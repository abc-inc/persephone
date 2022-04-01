package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	browser "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	persephone "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	shell "github.com/abc-inc/persephone/cmd/persephone/cmd/shell"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/hist"
	. "github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/persephone/playground"
	"github.com/abc-inc/persephone/types"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile = filepath.Join(Must(os.UserConfigDir()), "persephone", "config.yaml")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "persephone",
	Short: `A command line shell where you can execute Cypher against an instance of Neo4j. ` +
		`By default the shell is interactive but you can use it for scripting ` +
		`by passing cypher directly on the command line or by piping a file with cypher statements.`,
	Long:             ``,
	PersistentPreRun: connect,
	Run:              run,
	TraverseChildren: true,
}

var pass string
var lines []string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Args = cobra.MaximumNArgs(1)
	rootCmd.Flags().String("format", "auto", "Desired output format (default: auto).")
	rootCmd.Flags().StringSliceP("param", "P", nil, "Add a parameter to this session. Example: `-P \"number=3\"`. Can be specified multiple times.")
	rootCmd.Flags().Bool("version", false, "Print version of persephone and exit.")
	rootCmd.Flags().Bool("driver-version", false, "Print version of the Neo4j Driver used and exit.")
	rootCmd.Flags().StringP("file", "f", "", "Pass a file with cypher statements to be executed.")

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file ("+cfgFile+")")
	rootCmd.Flags().StringP("address", "a", "neo4j://localhost:7687", "address and port to connect to (default: neo4j://localhost:7687) (env: NEO4J_ADDRESS)")
	rootCmd.Flags().StringP("username", "u", "", "username to connect as (default: ). (env: NEO4J_USERNAME)")
	rootCmd.Flags().StringVarP(&pass, "password", "p", "", "password to connect with (default: ). (env: NEO4J_PASSWORD)")
	rootCmd.Flags().StringP("database", "d", "neo4j", "database to connect to (default: neo4j). (env: NEO4J_DATABASE)")

	rootCmd.AddCommand(
		browser.ClearCmd,
		browser.ConfigCmd,
		browser.QueriesCmd,
		browser.SchemaCmd,
		browser.SysinfoCmd,
		persephone.FormatCmd,
		persephone.IssueCmd,
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
	)

	MustNoErr(viper.BindPFlags(rootCmd.Flags()))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetEnvPrefix("NEO4J")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	end := Must(time.Parse("2006-01-02", "2022-04-08"))
	if time.Now().After(end) {
		color.Red("This early access preview of persephone is expired.")
	} else if end.Sub(time.Now()) < 3*24*time.Hour {
		color.Green("This early access preview of persephone will expire soon.")
	}

	log.SetFlags(0)
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func connect(cmd *cobra.Command, args []string) {
	if offline, ok := cmd.Annotations["offline"]; ok && offline == "true" {
		return
	}

	addr := viper.GetString("address")
	user := viper.GetString("username")
	pass := viper.GetString("password")
	db := viper.GetString("database")

	if user == "" && isatty.IsTerminal(os.Stdin.Fd()) {
		icons := func(set *survey.IconSet) {
			set.Question.Text = "Enter"
			set.Question.Format = ""
		}

		u := &survey.Input{Message: "username:", Default: "neo4j"}
		MustNoErr(survey.AskOne(u, &user, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
		p := &survey.Password{Message: "password:"}
		MustNoErr(survey.AskOne(p, &pass, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
	}

	fmt.Printf("Connecting to Neo4j database '%s' at '%s' as user '%s'.\n", db, addr, user)
	auth, user := graph.Auth(user + ":" + pass)
	conn := graph.NewConn(addr, user, auth, db)
	conn.DBName = db
}

func run(cmd *cobra.Command, args []string) {
	md, err := graph.GetConn().Metadata()
	if err != nil {
		log.Fatalln(err)
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
	history := hist.Load(histPath)
	defer func() { _ = history.Save() }()

	var p *prompt.Prompt
	p = prompt.New(func(cyp string) {
		if len(lines) == 0 && strings.HasPrefix(cyp, ":") {
			history.Add(cyp)
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

		history.Add(cyp)
		err := playground.Foo(graph.Request{Query: cyp, Params: graph.GetConn().Params})
		if err != nil {
			format.Writeln(err)
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
		res := es.GetCompletion(line, col, true)
		for _, i := range res.Items {
			if cyp == "" && (i.Type == types.Variable || i.Type == types.PropertyKey) {
				continue
			}
			pss = append(pss, prompt.Suggest{
				Text: i.Content,
			})
		}

		if len(pss) == 1 && strings.HasSuffix(cyp, pss[0].Text) {
			return nil
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
		prompt.OptionHistory(history.Entries()),
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
		format.Writeln(err)
	}
	return ""
}
