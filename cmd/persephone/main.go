// Copyright 2022 The persephone authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"

	browser "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	persephone "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	shell "github.com/abc-inc/persephone/cmd/persephone/cmd/shell"
	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/persephone/types"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/mattn/go-shellwords"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var p *prompt.Prompt
var e = editor.NewEditor("")
var compByConsCmd = make(map[string]repl.CompFunc)
var cfgFile = filepath.Join(internal.Must(os.UserConfigDir()), "persephone", "config.yaml")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "persephone",
	Short: `A command line shell where you can execute Cypher against an instance of Neo4j. ` +
		`By default the shell is interactive but you can use it for scripting ` +
		`by passing Cypher directly on the command line or by piping a file with Cypher statements.`,
	Long:             ``,
	PersistentPreRun: connect,
	Run:              run,
	TraverseChildren: true,
}

var lines []string

func init() {
	stdlog.SetFlags(0)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.TimeFieldFormat = "15:04:05"
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: zerolog.TimeFieldFormat,
	})

	cobra.OnInitialize(initConfig)

	rootCmd.Args = cobra.MaximumNArgs(1)
	rootCmd.Flags().Bool("driver-version", false, "Print version of the Neo4j Driver used and exit.")
	rootCmd.Flags().String("file", "", "Pass a file with cypher statements to be executed.")
	rootCmd.Flags().String("format", "auto", "Desired output format.")
	rootCmd.PersistentFlags().String("log-level", "info", "Level of details to be printed. (debug, info, error)")
	rootCmd.Flags().StringSliceP("param", "P", nil, "Add a parameter to this session. Example: `-P \"number=3\"`. Can be specified multiple times.")
	rootCmd.Flags().Bool("version", false, "Print version information and exit.")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file ("+cfgFile+")")
	rootCmd.PersistentFlags().StringP("address", "a", "neo4j://localhost:7687", "address and port to connect to (env: NEO4J_ADDRESS)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "username to connect as. (env: NEO4J_USERNAME)")
	rootCmd.PersistentFlags().StringP("database", "d", "neo4j", "database to connect to. (env: NEO4J_DATABASE)")

	rootCmd.AddCommand(
		browser.ChangePassCmd,
		browser.ClearCmd,
		browser.ConfigCmd,
		browser.QueriesCmd,
		browser.SchemaCmd,
		browser.SysinfoCmd,
		persephone.FormatCmd,
		persephone.IssueCmd,
		persephone.TemplateCmd,
		shell.BeginCmd,
		shell.CommitCmd,
		shell.ConnectCmd,
		shell.DisconnectCmd,
		shell.ExitCmd,
		shell.HelpCmd,
		shell.HistoryCmd,
		shell.ParamCmd,
		shell.ParamsCmd,
		shell.RollbackCmd,
		shell.SourceCmd,
		shell.UseCmd,
		DriverVersionCmd,
		VersionCmd,
	)

	internal.MustNoErr(viper.BindPFlag("address", rootCmd.Flag("address")))
	internal.MustNoErr(viper.BindPFlag("database", rootCmd.Flag("database")))
	internal.MustNoErr(viper.BindPFlag("format", rootCmd.Flag("format")))
	internal.MustNoErr(viper.BindPFlag("username", rootCmd.Flag("username")))

	if err := console.GetTmplMgr().Load(); err != nil {
		console.WriteErr(err)
	}

	var tmplComp = func(s string) (its []repl.Item) {
		for n := range console.GetTmplMgr().TmplsByPath {
			if strings.HasPrefix(n, s) {
				its = append(its, repl.Item{View: strings.TrimSuffix(n, console.TmplExt)})
			}
		}
		return
	}

	compByConsCmd[persephone.FQCmdName(persephone.FormatCmd)] = persephone.FormatComp

	compByConsCmd[persephone.FQCmdName(shell.HistoryCmd)] = repl.SubCmdComp(shell.HistoryCmd)

	compByConsCmd[persephone.FQCmdName(shell.SourceCmd)] = repl.PathComp

	compByConsCmd[persephone.FQCmdName(persephone.TemplateCmd)] = repl.SubCmdComp(persephone.TemplateCmd)
	compByConsCmd[persephone.FQCmdName(persephone.TemplateEditCmd)] = tmplComp
	compByConsCmd[persephone.FQCmdName(persephone.TemplateGetCmd)] = tmplComp
	compByConsCmd[persephone.FQCmdName(persephone.TemplateListCmd)] = repl.NoComp
	compByConsCmd[persephone.FQCmdName(persephone.TemplateSetCmd)] = tmplComp
	compByConsCmd[persephone.FQCmdName(persephone.TemplateWriteCmd)] = tmplComp

	compByConsCmd[persephone.FQCmdName(shell.UseCmd)] = shell.UseCompFunc()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetEnvPrefix("NEO4J")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug().Str("config", viper.ConfigFileUsed()).Msg("Loading config")
	}
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}

func connect(cmd *cobra.Command, args []string) {
	if offline, ok := cmd.Annotations["offline"]; ok && offline == "true" {
		return
	}
	for _, f := range []string{"driver-version", "help", "version"} {
		if f, err := cmd.Root().Flags().GetBool(f); err == nil && f {
			return
		}
	}
	if graph.IsConnected() {
		return
	}

	console.ChangeFmt(viper.GetString("format"))

	addr := viper.GetString("address")
	u := viper.GetString("username")
	p := viper.GetString("password")
	db := viper.GetString("database")

	if u == "" && isatty.IsTerminal(os.Stdin.Fd()) {
		u = console.Input("username:", "neo4j")
	}
	if p == "" && isatty.IsTerminal(os.Stdin.Fd()) {
		p = console.Pwd("password:")
	}

	log.Info().Str("db", db).Str("addr", addr).Str("user", u).
		Msg("Connecting to Neo4j database")

	auth, u := graph.Auth(u + ":" + p)
	conn := graph.NewConn(addr, u, auth, db)
	conn.DBName = db

	if isatty.IsTerminal(os.Stdin.Fd()) {
		consCmdCol := color.New(color.FgCyan).Sprint
		log.Info().Msgf("Type %s for a list of available commands or %s to exit the shell.",
			consCmdCol(":help"), consCmdCol(":exit"))
		log.Info().Msg("Note that Cypher queries must end with a semicolon.")
	}
}

func run(cmd *cobra.Command, args []string) {
	ll := strings.ToLower(internal.Must(cmd.Flags().GetString("log-level")))
	if l, err := zerolog.ParseLevel(ll); err == nil {
		zerolog.SetGlobalLevel(l)
	}

	if cmd == cmd.Root() && runRootCmd(cmd, args) {
		return
	}

	md, err := graph.GetConn().Metadata()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	ls := make([]string, len(md.Nodes))
	var pkeys []string
	for i, e := range md.Nodes {
		ls[i] = e.String()
		for _, p := range e.Properties {
			pkeys = append(pkeys, p)
		}
	}
	if len(pkeys) == 0 {
		pkeys = append(pkeys, md.Props...)
	}

	ts := make([]string, len(md.Rels))
	for i, r := range md.Rels {
		ts[i] = r.Type
		for p := range r.Properties {
			pkeys = append(pkeys, p)
		}
	}

	ccs := make([]graph.Cmd, len(cmd.Root().Commands()))
	for i, c := range cmd.Root().Commands() {
		ccs[i] = graph.Cmd{Name: c.Name(), Desc: strings.TrimPrefix(c.Name(), ":")}
	}

	schema := graph.Schema{
		Labels:   ls,
		RelTypes: ts,
		PropKeys: pkeys,
		Funcs:    md.Funcs,
		Procs:    md.Procs,
		ConCmds:  ccs,
	}

	e.SetSchema(schema)

	histPath := filepath.Join(internal.Must(os.UserCacheDir()), "persephone", "history")
	hist := repl.GetHistory()
	_ = hist.Load(histPath)
	defer func() {
		if err := hist.Save(histPath); err != nil {
			console.WriteErr(err)
		}
	}()

	p = prompt.New(func(cyp string) { executor(cyp, cmd) },
		completer,
		prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
			return breakline && in == "exit"
		}), prompt.OptionPrefix(""),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionCompletionWordSeparator(" "),
		prompt.OptionHistory(hist.Entries()),
		prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			if graph.GetConn().DBName == "" {
				return "Disconnected>", true
			}
			return fmt.Sprintf("%s@%s> ", graph.GetConn().Username(), graph.GetConn().DBName), len(lines) == 0
		}),
	)
	p.Run()
}

func runRootCmd(cmd *cobra.Command, args []string) bool {
	if internal.Must(cmd.Flags().GetBool("version")) {
		versionCmd()
		return true
	} else if internal.Must(cmd.Flags().GetBool("driver-version")) {
		driverVersionCmd()
		return true
	}

	if len(args) > 0 && !strings.HasPrefix(args[0], ":") {
		if err := console.Query(graph.Request{Query: strings.Join(args, " "), Params: graph.GetConn().Params}); err != nil {
			console.WriteErr(err)
		}
		return true
	}
	if !isatty.IsTerminal(os.Stdin.Fd()) {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			if err := console.Query(graph.Request{Query: sc.Text(), Params: graph.GetConn().Params}); err != nil {
				console.WriteErr(err)
			}
		}
		return true
	}
	return false
}

func executor(cyp string, cmd *cobra.Command) {
	hist := repl.GetHistory()
	if len(lines) == 0 && strings.HasPrefix(cyp, ":") {
		hist.Add(cyp)
		cyp = runConsCmd(cmd, cyp)
	}

	if cyp == "" {
		return
	}

	lines = append(lines, cyp)
	if !strings.HasSuffix(cyp, ";") {
		return
	}

	cyp = strings.Join(lines, "\n")
	lines = nil

	hist.Add(cyp)
	err := console.Query(graph.Request{Query: cyp, Params: graph.GetConn().Params})
	if err != nil {
		console.WriteErr(err)
	}
}

func completer(document prompt.Document) (ss []prompt.Suggest) {
	stmt := strings.TrimLeft(document.TextBeforeCursor(), " ")
	if stmt == "exit" || stmt == ":exit" {
		return nil
	}
	if stmt == "" || strings.IndexRune(");'\"", rune(stmt[len(stmt)-1])) >= 0 {
		return nil
	}

	buf := strings.Join(lines, "\n")
	buf += "\n" + stmt
	buf = strings.TrimPrefix(buf, "\n")

	var res comp.Result
	if strings.HasPrefix(stmt, ":") && strings.IndexByte(stmt, ' ') > 0 {
		res = compConsCmd(stmt)
	} else {
		e.Update(buf)
		line, col := editor.NewPosConv(buf).ToRelative(len(buf))
		res = e.GetCompletion(line, col, true)
	}

	for _, i := range res.Items {
		if stmt == "" && (i.Type == types.Variable || i.Type == types.PropertyKey) {
			continue
		}
		if strings.HasPrefix(i.View, "apoc.") && !strings.Contains(stmt, "apoc.") {
			continue
		}
		if i.View == strings.Trim(i.Content, "`") {
			ss = append(ss, prompt.Suggest{Text: i.View})
		} else {
			ss = append(ss, prompt.Suggest{Text: i.View, Description: i.Content})
		}
	}

	sep := " "
	start := res.Range.From.Col - 1
	if start >= 0 && start < len(document.CurrentLine()) {
		sep = document.CurrentLine()[start : start+1]
	}
	internal.MustNoErr(prompt.OptionCompletionWordSeparator(sep)(p))
	return ss
}

func compConsCmd(stmt string) (res comp.Result) {
	args, parts := "", strings.SplitN(stmt, " ", 3)
	var cmdComp repl.CompFunc
	for i := len(parts) - 1; i >= 0; i-- {
		if cmdComp = compByConsCmd[strings.Join(parts[:i], " ")]; cmdComp != nil {
			if len(parts) > i {
				args = strings.Join(parts[i:], " ")
			}
			parts = parts[:i]
			break
		}
	}

	if cmdComp != nil {
		for _, p := range cmdComp(args) {
			it := comp.Item{View: p.View, Content: p.Content}
			res.Items = append(res.Items, it)
		}

		start := strings.LastIndex(stmt, "/") + 1
		if start == 0 {
			start = len(parts[0]) + 1
		}
		res.Range = comp.Range{
			From: comp.LineCol{Line: 0, Col: start},
			To:   comp.LineCol{Line: 0, Col: len(stmt)},
		}
	}
	return
}

func runConsCmd(cmd *cobra.Command, stmt string) string {
	args := internal.Must(shellwords.Parse(stmt))
	cmd.Root().SetArgs(args)
	if args[0] == ":param" {
		args = strings.SplitN(stmt, " ", 3)
		cmd.Root().SetArgs(args)
	}
	if err := cmd.Execute(); err != nil {
		console.WriteErr(err)
	}
	return ""
}
