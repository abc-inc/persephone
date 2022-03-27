package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abc-inc/persephone/format"
	"github.com/abc-inc/persephone/graph"
	. "github.com/abc-inc/persephone/internal"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConnectCmd = &cobra.Command{
	Use:   ":connect",
	Short: "Connects to a database",
	Run:   connectCmd,
}

func init() {
	ConnectCmd.Flags().StringP("username", "u", "", "username to connect as (default: ). (env: NEO4J_USERNAME)")
	ConnectCmd.Flags().StringP("password", "p", "", "password to connect with (default: ). (env: NEO4J_PASSWORD)")
	ConnectCmd.Flags().StringP("database", "d", "neo4j", "database to connect to (default: neo4j). (env: NEO4J_DATABASE)")
}

func connectCmd(cmd *cobra.Command, args []string) {
	if graph.GetConn() != nil && graph.GetConn().Driver != nil {
		format.Writeln("Already connected")
		return
	}

	icons := func(set *survey.IconSet) {
		set.Question.Text = "Enter"
		set.Question.Format = ""
	}

	var user, pass string
	addr := viper.GetString("address")
	db := viper.GetString("database")

	u := &survey.Input{Message: "username:", Default: "neo4j"}
	MustNoErr(survey.AskOne(u, &user, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
	p := &survey.Password{Message: "password:"}
	MustNoErr(survey.AskOne(p, &pass, survey.WithValidator(survey.Required), survey.WithIcons(icons)))

	fmt.Printf("Connecting to Neo4j database '%s' at '%s' as user '%s'.\n", db, addr, user)
	driver := Must(neo4j.NewDriver(addr, neo4j.BasicAuth(user, pass, "")))
	conn := graph.NewConn(driver, db)
	conn.DBName = db
}
