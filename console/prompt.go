package console

import (
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abc-inc/persephone/internal"
	"github.com/briandowns/spinner"
)

func Input(msg string, def string) (res string) {
	in := &survey.Input{Message: msg, Default: def}
	internal.MustNoErr(survey.AskOne(in, &res, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
	return
}

func Pwd(msg string) (res string) {
	in := &survey.Password{Message: msg}
	internal.MustNoErr(survey.AskOne(in, &res, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
	return
}

func icons(set *survey.IconSet) {
	set.Question.Text = "Enter"
	set.Question.Format = ""
}

func NewSpinner() *spinner.Spinner {
	dotStyle := spinner.CharSets[11]
	return spinner.New(dotStyle, 120*time.Millisecond,
		spinner.WithWriter(os.Stderr),
		spinner.WithColor("fgCyan"),
		spinner.WithFinalMSG("\033[2K\r"))
}
