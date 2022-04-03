package console

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/abc-inc/persephone/internal"
)

func Icons(set *survey.IconSet) {
	set.Question.Text = "Enter"
	set.Question.Format = ""
}

func Input(msg string, def string) (res string) {
	in := &survey.Input{Message: msg}
	internal.MustNoErr(survey.AskOne(in, &res, survey.WithValidator(survey.Required), survey.WithIcons(Icons)))
	return
}

func Pwd(msg string) (res string) {
	in := &survey.Password{Message: msg}
	internal.MustNoErr(survey.AskOne(in, &res, survey.WithValidator(survey.Required), survey.WithIcons(Icons)))
	return
}
