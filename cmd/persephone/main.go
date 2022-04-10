package main

import (
	"bufio"
	"fmt"
	stdlog "log"
	"os"
	"path/filepath"
	"strings"
	"time"

	browser "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	persephone "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	shell "github.com/abc-inc/persephone/cmd/persephone/cmd/shell"
	"github.com/abc-inc/persephone/comp"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
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

var compByConsCmd = make(map[string]repl.CompFunc)

var cfgFile = filepath.Join(Must(os.UserConfigDir()), "persephone", "config.yaml")

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
	rootCmd.Flags().String("log-level", "info", "Level of details to be printed. (debug, info, error)")
	rootCmd.Flags().StringSliceP("param", "P", nil, "Add a parameter to this session. Example: `-P \"number=3\"`. Can be specified multiple times.")
	rootCmd.Flags().Bool("version", false, "Print version information and exit.")

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file ("+cfgFile+")")
	rootCmd.Flags().StringP("address", "a", "neo4j://localhost:7687", "address and port to connect to (env: NEO4J_ADDRESS)")
	rootCmd.Flags().StringP("username", "u", "", "username to connect as. (env: NEO4J_USERNAME)")
	rootCmd.Flags().StringP("password", "p", "", "password to connect with. (env: NEO4J_PASSWORD)")
	rootCmd.Flags().StringP("database", "d", "neo4j", "database to connect to. (env: NEO4J_DATABASE)")

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

	MustNoErr(viper.BindPFlag("address", rootCmd.Flag("address")))
	MustNoErr(viper.BindPFlag("database", rootCmd.Flag("database")))
	MustNoErr(viper.BindPFlag("format", rootCmd.Flag("format")))
	MustNoErr(viper.BindPFlag("username", rootCmd.Flag("username")))

	compByConsCmd[shell.SourceCmd.Name()] = repl.PathComp
	compByConsCmd[persephone.TemplateCmd.Name()] =
		func(str string) (its []repl.Item) {
			for n, t := range console.Tmpls {
				its = append(its, repl.Item{View: n, Content: t.Root.String()})
			}
			return
		}
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
	end := Must(time.Parse("2006-01-02", "2022-04-08"))
	if time.Now().After(end) {
		color.Red("This early access preview of persephone is expired.")
	} else if end.Sub(time.Now()) < 3*24*time.Hour {
		color.Green("This early access preview of persephone will expire soon.")
	}

	cobra.CheckErr(rootCmd.Execute())
}

func connect(cmd *cobra.Command, args []string) {
	if offline, ok := cmd.Annotations["offline"]; ok && offline == "true" {
		return
	}
	for _, f := range []string{"driver-version", "help", "version"} {
		if Must(cmd.Root().Flags().GetBool(f)) {
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
	ll := strings.ToLower(Must(cmd.Flags().GetString("log-level")))
	if l, err := zerolog.ParseLevel(ll); err == nil {
		zerolog.SetGlobalLevel(l)
	}

	if Must(cmd.Flags().GetBool("version")) {
		versionCmd()
		return
	} else if Must(cmd.Flags().GetBool("driver-version")) {
		driverVersionCmd()
		return
	}

	if cmd == cmd.Root() && len(args) > 0 && !strings.HasPrefix(args[0], ":") {
		if err := query(graph.Request{strings.Join(args, " "), graph.GetConn().Params}); err != nil {
			console.Writeln(err)
		}
		return
	}
	if cmd == cmd.Root() && !isatty.IsTerminal(os.Stdin.Fd()) {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			if err := query(graph.Request{sc.Text(), graph.GetConn().Params}); err != nil {
				console.Writeln(err)
			}
		}
		return
	}

	md, err := graph.GetConn().Metadata()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	var ls, ts, pkeys []string
	for _, e := range md.Nodes {
		ls = append(ls, e.String())
		for _, p := range e.Properties {
			pkeys = append(pkeys, p)
		}
	}
	if len(pkeys) == 0 {
		pkeys = append(pkeys, md.Props...)
	}

	for _, r := range md.Rels {
		ts = append(ts, r.Type)
		for p := range r.Properties {
			pkeys = append(pkeys, p)
		}
	}

	var ccs []graph.Cmd
	for _, c := range cmd.Root().Commands() {
		ccs = append(ccs, graph.Cmd{
			Name: c.Name(),
			Desc: strings.TrimPrefix(c.Name(), ":"),
		})
	}

	schema := graph.Schema{
		Labels:   ls,
		RelTypes: ts,
		PropKeys: pkeys,
		Funcs:    md.Funcs,
		Procs:    md.Procs,
		ConCmds:  ccs,
	}

	es := editor.NewEditorSupport("")
	es.SetSchema(schema)

	histPath := filepath.Join(Must(os.UserCacheDir()), "persephone", "history")
	hist := repl.GetHistory()
	hist.Load(histPath)
	defer func() {
		if err := hist.Save(histPath); err != nil {
			console.Writeln(err)
		}
	}()

	var p *prompt.Prompt
	p = prompt.New(func(cyp string) {
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
		err := query(graph.Request{
			Query:  cyp,
			Params: graph.GetConn().Params,
		})
		if err != nil {
			console.Writeln(err)
		}
	}, func(document prompt.Document) (pss []prompt.Suggest) {
		cyp := document.TextBeforeCursor()
		if cyp == "exit" || cyp == ":exit" {
			return nil
		}
		if cyp == "" || strings.IndexRune(");'\"", rune(cyp[len(cyp)-1])) >= 0 {
			return nil
		}

		buf := strings.Join(lines, "\n")
		buf += "\n" + cyp
		buf = strings.TrimPrefix(buf, "\n")

		es.Update(buf)
		line, col := editor.NewPosConv(buf).ToRelative(len(buf))
		var res comp.Result

		parts := strings.SplitN(cyp, " ", 2)
		if cmdComp, ok := compByConsCmd[parts[0]]; ok && len(parts) > 1 {
			res = comp.Result{}
			for _, p := range cmdComp(parts[1]) {
				it := comp.Item{View: p.View, Content: p.Content}
				res.Items = append(res.Items, it)
			}

			start := strings.LastIndex(cyp, "/") + 1
			if start == 0 {
				start = len(parts[0]) + 1
			}
			res.Range = comp.Range{
				From: comp.LineCol{Line: 0, Col: start},
				To:   comp.LineCol{Line: 0, Col: len(cyp)},
			}
		} else {
			res = es.GetCompletion(line, col, true)
		}
		for _, i := range res.Items {
			if cyp == "" && (i.Type == types.Variable || i.Type == types.PropertyKey) {
				continue
			}
			if strings.HasPrefix(i.View, "apoc.") && !strings.Contains(cyp, "apoc.") {
				continue
			}
			if i.View == strings.Trim(i.Content, "`") {
				pss = append(pss, prompt.Suggest{Text: i.View})
			} else {
				pss = append(pss, prompt.Suggest{Text: i.View, Description: i.Content})
			}
		}

		sep := " "
		start := res.Range.From.Col - 1
		if start >= 0 && start < len(document.CurrentLine()) {
			sep = document.CurrentLine()[start : start+1]
		}
		MustNoErr(prompt.OptionCompletionWordSeparator(sep)(p))
		return
	}, prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
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

func runConsCmd(cc *cobra.Command, cyp string) string {
	args := Must(shellwords.Parse(cyp))
	cc.Root().SetArgs(args)
	if args[0] == ":param" {
		args = strings.SplitN(cyp, " ", 3)
		cc.Root().SetArgs(args)
	}
	if err := cc.Execute(); err != nil {
		console.Writeln(err)
	}
	return ""
}

func query(req graph.Request) error {
	log.Debug().Str("statement", req.Query).Fields(req.Params).Msg("Executing")

	if console.FormatName() == "raw" || console.FormatName() == "rawc" {
		return queryRaw(req)
	}
	return queryResult(req)
}

func queryRaw(req graph.Request) error {
	sp := console.NewSpinner()
	sp.Start()

	t := graph.NewTypedTemplate[map[string]interface{}](graph.GetConn())
	ms, sum, err := t.Query(req.Query, req.Params, graph.NewRawResultRowMapper())

	sp.Stop()
	if err == nil {
		console.Writeln(ms)
		console.WriteSummary(len(ms), sum)
	}
	return err
}

func queryResult(req graph.Request) error {
	sp := console.NewSpinner()
	sp.Start()

	t := graph.NewTypedTemplate[graph.Result](graph.GetConn())
	rs, sum, err := t.Query(req.Query, req.Params, graph.NewResultRowMapper())

	sp.Stop()
	if err == nil {
		err = console.WriteResult(rs, sum)
		console.WriteSummary(len(rs), sum)
	}
	return err
}
