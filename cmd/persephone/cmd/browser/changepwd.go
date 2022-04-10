package cmd

import (
	"errors"
	"os"

	"github.com/abc-inc/persephone/console"
	"github.com/abc-inc/persephone/graph"
	"github.com/abc-inc/persephone/internal"
	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const cypSetPass = "ALTER CURRENT USER SET PASSWORD FROM $old TO $new"

var errPwdEmpty = errors.New("current and new password must be provided")
var errPwdMismatch = errors.New("the two entered passwords must be the same")

var ChangePassCmd = &cobra.Command{
	Use:    ":change-password",
	Short:  "Change the user password",
	Run:    changePassCmd,
	Hidden: true,
}

func init() {
	ChangePassCmd.Flags().StringP("new-password", "n", "", "new password to connect with")
	ChangePassCmd.Flags().StringP("password", "p", "", "current password")
}

func changePassCmd(cmd *cobra.Command, args []string) {
	p := internal.Must(cmd.Flags().GetString("password"))
	newP := internal.Must(cmd.Flags().GetString("new-password"))
	ChangePass(p, newP, newP)
}

func ChangePass(p, newP1, newP2 string) {
	if isatty.IsTerminal(os.Stdin.Fd()) {
		if p == "" {
			p = console.Pwd("password:")
		}
		if newP1 == "" {
			newP1 = console.Pwd("new password:")
			newP2 = console.Pwd("new password (repeat):")
		}
	}

	if p == "" || newP1 == "" {
		console.Writeln(errPwdEmpty)
		return
	}
	if newP1 != newP2 {
		console.Writeln(errPwdMismatch)
		return
	}

	_, err := graph.GetConn().Session().WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		return tx.Run(cypSetPass, map[string]interface{}{"old": p, "new": newP1})
	})
	if err != nil {
		console.Writeln(err)
	} else {
		log.Info().Msg(color.GreenString("Password changed successfully"))
	}
}
