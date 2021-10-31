package cypher

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"regexp"
	"strings"
)

type Mode int

const (
	Keyword Mode = iota
	Node
	Variable
)

func (m Mode) Separator() string {
	return []string{" ", ":", "("}[m]
}

type Prompt struct {
	prompt      *prompt.Prompt
	complByMode map[Mode]func(prompt.Document) []prompt.Suggest
}

func NewPrompt() *Prompt {
	executor := func(in string) {}
	p := &Prompt{}
	p.complByMode = make(map[Mode]func(document prompt.Document) []prompt.Suggest)
	p.prompt = prompt.New(executor, p.completer)
	return p
}

func (p *Prompt) Switch(m Mode) func(prompt.Document) []prompt.Suggest {
	opt := prompt.OptionCompletionWordSeparator(m.Separator())
	opt(p.prompt)
	if compl, ok := p.complByMode[m]; ok {
		return compl
	}
	return defaultCompl
}

func (p *Prompt) Register(m Mode, compl func(prompt.Document) []prompt.Suggest) {
	p.complByMode[m] = compl
}

func (p *Prompt) Run() {
	p.prompt.Run()
}

func (p *Prompt) Input() string {
	return p.prompt.Input()
}

func (p *Prompt) completer(d prompt.Document) []prompt.Suggest {
	vars := findVars(d.TextBeforeCursor())
	fmt.Println(vars)

	p.Switch(Keyword)
	line := d.GetWordBeforeCursor()
	compl := defaultCompl
	if strings.HasSuffix(line, ":") {
		compl = p.Switch(Node)
	}

	return compl(d)
}

func findVars(s string) []string {
	as := `\s+AS\s+(?P<var3>[^\s]+)`
	rel := `\[(?P<var1>[^]:]+)(:[^]]+)?\]`
	node := `\((?P<var2>[^):]+)(:[^)]+)?\)`
	anonNode := `\((?P<var2>[^):]+)?(:[^)]+)?\)`
	pattern := fmt.Sprintf("(-%s-[<>]?%s)|([\\s-]%s)|(%s)", rel, anonNode, node, as)
	re := regexp.MustCompile(pattern)
	fmt.Println(s)
	subs := re.FindAllStringSubmatch(s, -1)
	vars := []string{}
	for _, sub := range subs {
		fmt.Println(">>>>"+strings.Join(sub, "^"))
		for i, n := range re.SubexpNames() {
			if strings.HasPrefix(n, "var") && sub[i] != "" {
				if sub[i] != "" {
					vars = append(vars, sub[i])
				}
			}
		}
	}
	return vars
}

func defaultCompl(_ prompt.Document) []prompt.Suggest {
	return nil
}
