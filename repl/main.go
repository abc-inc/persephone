package repl

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func NewRepl() {
	p := prompt.New(func(s string) {
		fmt.Println("EXEC ", s)
	}, func(document prompt.Document) (ps []prompt.Suggest) {
		s := document.TextBeforeCursor()
		is := Foo(s)
		for _, i := range is {
			ps = append(ps, prompt.Suggest{
				Text:        i.View,
				Description: "",
			})
		}
		return
	})

	p.Run()
}

func main()  {
	NewRepl()
}
