package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	browser "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	persephone "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	shell "github.com/abc-inc/persephone/cmd/persephone/cmd/shell"
	"github.com/abc-inc/persephone/editor"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/persephone/playground"
	"github.com/abc-inc/persephone/types"
	"github.com/c-bata/go-prompt"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
	Long: ``,
	Run:  run,
}

var pass string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Args = cobra.MaximumNArgs(1)
	rootCmd.Flags().String("format", "auto", "Desired output format (default: auto).")
	rootCmd.Flags().StringSliceP("param", "P", nil, "Add a parameter to this session. Example: `-P \"number=3\"`. Can be specified multiple times.")
	rootCmd.Flags().BoolP("version", "v", false, "Print version of persephone and exit.")
	rootCmd.Flags().Bool("driver-version", false, "Print version of the Neo4j Driver used and exit.")
	rootCmd.Flags().StringP("file", "f", "", "Pass a file with cypher statements to be executed.")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file ("+cfgFile+")")
	rootCmd.PersistentFlags().StringP("address", "a", "neo4j://localhost:7687", "address and port to connect to (default: neo4j://localhost:7687) (env: NEO4J_ADDRESS)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "username to connect as (default: ). (env: NEO4J_USERNAME)")
	rootCmd.PersistentFlags().StringVarP(&pass, "password", "p", "", "password to connect with (default: ). (env: NEO4J_PASSWORD)")
	rootCmd.PersistentFlags().StringP("database", "d", "neo4j", "database to connect to (default: neo4j). (env: NEO4J_DATABASE)")

	rootCmd.AddCommand(
		browser.ClearCmd,
		browser.ConfigCmd,
		browser.QueriesCmd,
		browser.SchemaCmd,
		browser.SysinfoCmd,
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

	MustNoErr(viper.BindPFlags(rootCmd.PersistentFlags()))
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
	Execute()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func run(cmd *cobra.Command, args []string) {
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
	driver := Must(neo4j.NewDriver(addr, neo4j.BasicAuth(user, pass, "")))
	conn := graph.NewConn(driver)
	conn.DBName = db

	md, err := conn.Metadata()
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

	schema := graph.Schema{
		Labels:   ls,
		RelTypes: ts,
		PropKeys: pkeys,
		Funcs:    md.Funcs,
		Procs:    md.Procs,
	}

	es := editor.NewEditorSupport("")
	es.SetSchema(schema)

	var p *prompt.Prompt
	p = prompt.New(func(cyp string) {
		if len(lines) == 0 && strings.HasPrefix(cyp, ":") {
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

		err := playground.Foo(os.Stdout, conn, graph.Request{Query: cyp})
		if err != nil {
			log.Fatalln(err)
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

		return
	}, prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
		return breakline && in == "exit"
	}), prompt.OptionPrefix(fmt.Sprintf("%s@%s> ", user, db)),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionCompletionWordSeparator(" :(."),
		prompt.OptionLivePrefix(func() (prefix string, useLivePrefix bool) {
			return "", len(lines) > 0
		}),
	)
	p.Run()
}

func runConsCmd(cc *cobra.Command, cyp string) string {
	cmd, arg, _ := strings.Cut(cyp, " ")
	switch cmd {
	case ":begin":
		break
	case ":clear":
		browser.ClearCmd.Run(cc, strings.Fields(cyp))
		break
	case ":config":
		break
	case "format":
		break
	case ":help":
		shell.HelpCmd.Run(cc, strings.Fields(cyp))
		break
	case ":history":
		break
	case ":queries":
		break
	case ":param":
		r := regexp.MustCompile("\\s*=>\\s*")
		kv := r.Split(arg, 2)
		var v map[string]interface{}
		MustNoErr(json.Unmarshal([]byte(kv[1]), &v))
		params[kv[0]] = v
		break
	case ":params":
		j := Must(json.MarshalIndent(params, "", "  "))
		fmt.Println(string(j))
		break
	case ":schema":
		return "CALL db.indexes() YIELD name AS `Index Name`, type AS Type, uniqueness AS Uniqueness, " +
			"entityType AS EntityType, labelsOrTypes AS LabelsOrTypes, properties AS Properties, state AS State " +
			"ORDER BY name;"
	case ":server":
		break
	case ":style":
		break
	case "sysinfo":
		break
	default:
		break
	}
	return ""
}

func historyCmd(cmd *cobra.Command, args []string) {
	fmt.Println("history")
}

var history []string
var lines []string
var params map[string]interface{}
