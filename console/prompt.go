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

package console

import (
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/abc-inc/persephone/internal"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

// Input prompts the user to input a single message.
func Input(msg string, def string) (res string) {
	in := &survey.Input{Message: msg, Default: def}
	internal.MustNoErr(survey.AskOne(in, &res, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
	return
}

// Pwd is like Input but does not echo the input.
func Pwd(msg string) (res string) {
	in := &survey.Password{Message: msg}
	internal.MustNoErr(survey.AskOne(in, &res, survey.WithValidator(survey.Required), survey.WithIcons(icons)))
	return
}

// icons configures the IconSet used when prompting the user for input.
func icons(set *survey.IconSet) {
	set.Question.Text = "Enter"
	set.Question.Format = ""
}

// NewSpinner creates a new default Spinner, which writes to stderr.
// Additional options can be set before it is started.
func NewSpinner() *spinner.Spinner {
	dotStyle := spinner.CharSets[11]
	return spinner.New(dotStyle, 120*time.Millisecond,
		spinner.WithWriter(os.Stderr),
		spinner.WithColor("fgCyan"),
		spinner.WithFinalMSG("\033[2K\r"))
}

func SuccessIcon() string {
	return color.GreenString("✓")
}

func WarningIcon() string {
	return color.YellowString("!")
}

func FailureIcon() string {
	return color.RedString("X")
}
