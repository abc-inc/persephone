// Copyright 2022 The Persephone authors
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

package cmd

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	browser "github.com/abc-inc/persephone/cmd/persephone/cmd/browser"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/cmdutil"
	"github.com/abc-inc/persephone/cmd/persephone/cmd/help"
	persephone "github.com/abc-inc/persephone/cmd/persephone/cmd/persephone"
	shell "github.com/abc-inc/persephone/cmd/persephone/cmd/shell"
	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/console/repl"
	"github.com/abc-inc/persephone/internal"
	"github.com/abc-inc/roland/graph"
	"github.com/mattn/go-isatty"
	"github.com/mattn/go-shellwords"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	lines         []string
	compByConsCmd = make(map[string]repl.CompFunc)
)

func NewCmdRoot(f *cmdutil.Factory) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "persephone <command> [flags]",
		Short: "Neo4j database client",
		Long: `A command line shell where you can execute Cypher against an instance of Neo4j. ` +
			`By default the shell is interactive but you can use it for scripting ` +
			`by passing Cypher directly on the command line or by piping a file with Cypher statements.`,
		Example: heredoc.Doc(`
			$ persephone "MATCH (n) RETURN count(n);"
			$ persephone --address neo4j://localhost:7687 --database system --username neo4j
			$ persephone :sysinfo
			$ persephone :help environment`),
		Args: cobra.MaximumNArgs(1),
		Annotations: map[string]string{
			"help:feedback": heredoc.Doc(`
				Open an issue using 'persephone issue'
			`),
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmdutil.IsAuthCheckEnabled(cmd) && !graph.IsConnected() {
				cmdutil.Connect(f.SessionConfig())
			}
		},
		Run:               func(cmd *cobra.Command, args []string) { run(f, cmd, args) },
		PersistentPostRun: cmdutil.ResetFlags,
	}

	rootCmd.Flags().Bool("driver-version", false, "Print version of the Neo4j Driver used and exit")
	rootCmd.Flags().String("file", "", "Pass a file with cypher statements to be executed")
	rootCmd.PersistentFlags().String("format", "", "Desired output format")
	rootCmd.PersistentFlags().String("log-level", "info", "Level of details to be printed (debug, info, error)")
	rootCmd.Flags().StringArrayP("param", "P", nil, "Add a parameter to this session (can be specified multiple times). Example: `-P \"number=3\"`")
	rootCmd.PersistentFlags().Bool("help", false, "Show help and exit")
	rootCmd.Flags().Bool("version", false, "Print version information and exit")

	rootCmd.PersistentFlags().StringVarP(f.SessionConfig().CfgFile, "config", "c", *f.SessionConfig().CfgFile, "Config file to use")
	rootCmd.PersistentFlags().StringVarP(f.SessionConfig().Address, "address", "a", *f.SessionConfig().Address, "Address and port to connect to (env: NEO4J_ADDRESS)")
	rootCmd.PersistentFlags().StringVarP(f.SessionConfig().Username, "username", "u", *f.SessionConfig().Username, "Username to connect as (env: NEO4J_USERNAME)")
	rootCmd.PersistentFlags().StringVarP(f.SessionConfig().Database, "database", "d", *f.SessionConfig().Database, "Database to connect to (env: NEO4J_DATABASE)")

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		rootHelpFunc(f, cmd, args)
	})
	// rootCmd.SetUsageFunc(rootUsageFunc)

	// Child commands
	cmdColors := persephone.NewCmdColors(f)
	cmdFormat := persephone.NewCmdFormat(f)
	cmdHistory := shell.NewCmdHistory(f)
	cmdSource := shell.NewCmdSource(f)
	cmdTemplate := persephone.NewCmdTemplate(f)
	cmdUse := shell.NewCmdUse(f)

	rootCmd.AddCommand(
		browser.NewCmdChangePass(f),
		browser.NewCmdClear(f),
		browser.NewCmdConfig(f),
		browser.NewCmdDBs(f),
		browser.NewCmdQueries(f),
		browser.NewCmdSchema(f),
		browser.NewCmdStatus(f),
		browser.NewCmdSysInfo(f),
		cmdColors,
		cmdFormat,
		persephone.NewCmdIssue(f),
		cmdTemplate,
		shell.NewCmdBegin(f),
		shell.NewCmdCommit(f),
		shell.NewCmdConnect(f),
		shell.NewCmdDisconnect(f),
		shell.NewCmdExit(f),
		shell.NewCmdHelp(f),
		cmdHistory,
		shell.NewCmdParam(f),
		shell.NewCmdRollback(f),
		cmdSource,
		cmdUse,
		NewCmdDriverVersion(f),
		NewCmdVersion(f),
	)

	// Completion
	if err := console.GetTmplMgr().Load(); err != nil {
		console.WriteErr(err)
	}
	tmplComp := func(s string) (its []repl.Item) {
		for tp := range console.GetTmplMgr().TmplsByPath {
			tp = strings.TrimSuffix(filepath.Base(tp), console.TmplExt)
			if strings.HasPrefix(tp, s) {
				its = append(its, repl.Item{View: tp})
			}
		}
		return
	}

	compByConsCmd[persephone.FQCmdName(cmdColors)] = persephone.ColorsComp

	compByConsCmd[persephone.FQCmdName(cmdFormat)] = persephone.FormatComp

	compByConsCmd[persephone.FQCmdName(cmdHistory)] = repl.SubCmdComp(cmdHistory)

	compByConsCmd[persephone.FQCmdName(cmdSource)] = repl.PathComp

	compByConsCmd[persephone.FQCmdName(cmdTemplate)] = repl.SubCmdComp(cmdTemplate)
	compByConsCmd[persephone.FQCmdName(persephone.FindCmd(cmdTemplate, "edit"))] = tmplComp
	compByConsCmd[persephone.FQCmdName(persephone.FindCmd(cmdTemplate, "get"))] = tmplComp
	compByConsCmd[persephone.FQCmdName(persephone.FindCmd(cmdTemplate, "list"))] = repl.NoComp
	compByConsCmd[persephone.FQCmdName(persephone.FindCmd(cmdTemplate, "set"))] = tmplComp
	compByConsCmd[persephone.FQCmdName(persephone.FindCmd(cmdTemplate, "write"))] = tmplComp

	compByConsCmd[persephone.FQCmdName(cmdUse)] = shell.UseCompFunc()

	// Help topics
	for _, t := range help.Topics {
		rootCmd.AddCommand(help.NewCmdHelpTopic(t))
	}
	return rootCmd
}

func run(f *cmdutil.Factory, cmd *cobra.Command, args []string) {
	if cmd == cmd.Root() && runRootCmd(cmd, args) {
		return
	}
	runRepl(f.Config(), cmd)
}

func runRootCmd(cmd *cobra.Command, args []string) bool {
	if internal.Must(cmd.Flags().GetBool("version")) {
		appVersion()
		return true
	} else if internal.Must(cmd.Flags().GetBool("driver-version")) {
		driverVersion()
		return true
	} else if ps := internal.Must(cmd.Flags().GetStringArray("param")); len(ps) > 0 {
		for _, p := range ps {
			k, v, ok := strings.Cut(p, "=")
			if !ok {
				internal.MustNoErr(cmd.Usage())
				return true
			}
			if shell.SetParam(k, v) != nil {
				return true
			}
		}
	}

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		hist := repl.NewHistory(1)
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			log.Info().Str("statement", sc.Text()).Msg("Executing")
			if err := executor(sc.Text(), cmd, hist); err != nil {
				console.WriteErr(err)
				return true
			}
		}
		return true
	} else if len(args) == 0 {
		return false
	}

	if strings.HasPrefix(args[0], ":") {
		consCmd, args2, err := cmd.Root().Find(args)
		if err != nil {
			console.WriteErr(err)
		} else if cmd == consCmd {
			console.WriteErr(cmd.Usage())
		} else {
			console.WriteErr(runConsCmd(consCmd, strings.Join(args2, " ")))
		}
	} else {
		r := graph.Request{Query: strings.Join(args, " "), Params: graph.GetConn().Params}
		if err := console.Query(r); err != nil {
			console.WriteErr(err)
		}
	}
	return true
}

func executor(cyp string, cmd *cobra.Command, hist *repl.History) error {
	if cyp == "" {
		return nil
	}

	if len(lines) == 0 && strings.HasPrefix(cyp, ":") {
		hist.Add(cyp)
		return runConsCmd(cmd, cyp)
	}

	lines = append(lines, cyp)
	if !strings.HasSuffix(cyp, ";") {
		return nil
	}

	cyp = strings.Join(lines, "\n")
	lines = nil

	hist.Add(cyp)
	r := graph.Request{Query: cyp, Params: graph.GetConn().Params}
	if err := console.Query(r); err != nil {
		return err
	}
	return nil
}

func runConsCmd(cmd *cobra.Command, stmt string) error {
	args := internal.Must(shellwords.Parse(stmt))
	cmd.Root().SetArgs(args)
	if len(args) == 4 && args[0] == ":param" && args[1] == "set" {
		args = strings.SplitN(stmt, " ", 4)
		cmd.Root().SetArgs(args)
	}
	err := cmd.Execute()
	cmdutil.ResetFlags(cmd, args)
	return err
}
