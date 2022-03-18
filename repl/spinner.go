package repl

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
)

func NewSpinner() *spinner.Spinner {
	dotStyle := spinner.CharSets[11]
	return spinner.New(dotStyle, 120*time.Millisecond,
		spinner.WithWriter(os.Stderr),
		spinner.WithColor("fgCyan"),
		spinner.WithFinalMSG("\033[2K\r"))
}
