package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/abc-inc/merovingian/cypher"
	"github.com/abc-inc/merovingian/ndb"
	"github.com/abc-inc/merovingian/playground"
	"github.com/abc-inc/merovingian/types"
	"github.com/c-bata/go-prompt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile = filepath.Join(must(os.UserConfigDir()), "merovingian", "config.yaml")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "merovingian",
	Short: "",
	Long:  ``,
	Run:   run,
}

var pass string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Args = cobra.MaximumNArgs(1)
	rootCmd.Flags().String("format", "auto", "Desired output format (default: auto).")
	rootCmd.Flags().StringSliceP("param", "P", nil, "Add a parameter to this session. Example: `-P \"number=3\"`. Can be specified multiple times.")
	rootCmd.Flags().BoolP("version", "v", false, "Print version of merovingian and exit.")
	rootCmd.Flags().Bool("driver-version", false, "Print version of the Neo4j Driver used and exit.")
	rootCmd.Flags().StringP("file", "f", "", "Pass a file with cypher statements to be executed.")

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", cfgFile, "config file ("+cfgFile+")")
	rootCmd.PersistentFlags().StringP("address", "a", "neo4j://localhost:7687", "address and port to connect to (default: neo4j://localhost:7687) (env: NEO4J_ADDRESS)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "username to connect as (default: ). (env: NEO4J_USERNAME)")
	rootCmd.PersistentFlags().StringVarP(&pass, "password", "p", "", "password to connect with (default: ). (env: NEO4J_PASSWORD)")
	rootCmd.PersistentFlags().StringP("database", "d", "", "database to connect to (default: ). (env: NEO4J_DATABASE)")

	mustNoErr(viper.BindPFlags(rootCmd.PersistentFlags()))
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

	driver := must(neo4j.NewDriver(addr, neo4j.BasicAuth(user, pass, "")))
	conn := ndb.NewConn(driver)
	conn.DBName = db

	ns, rs, err := conn.Metadata()
	if err != nil {
		log.Fatalln(err)
	}

	var ls, ts, pkeys []string
	for _, e := range ns {
		ls = append(ls, e.String())
		for _, p := range e.Properties {
			pkeys = append(pkeys, p)
		}
	}
	for _, r := range rs {
		ts = append(ts, r.Type)
		for p := range r.Properties {
			pkeys = append(pkeys, p)
		}
	}

	schema := ndb.Schema{
		Labels:   ls,
		RelTypes: ts,
		PropKeys: pkeys,
	}

	es := cypher.NewEditorSupport("")
	es.SetSchema(schema)

	var p *prompt.Prompt
	p = prompt.New(func(cyp string) {
		err := playground.Foo(os.Stdout, conn, ndb.Request{Query: cyp})
		if err != nil {
			log.Fatalln(err)
		}
	}, func(document prompt.Document) (pss []prompt.Suggest) {
		cyp := document.TextBeforeCursor()
		if cyp == "exit" || cyp == ":exit" {
			return nil
		}
		if cyp != "" && strings.IndexRune(" );'\"", rune(cyp[len(cyp)-1])) >= 0 {
			return nil
		}

		es.Update(cyp)
		line, col := cypher.NewPosConv(cyp).ToRelative(len(cyp))
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
		prompt.OptionCompletionWordSeparator(" :(."))
	p.Run()
}

func mustNoErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func must[T any](a T, err error) T {
	mustNoErr(err)
	return a
}
