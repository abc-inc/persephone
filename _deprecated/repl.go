package _deprecated

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/abc-inc/merovingian/comp"
	"github.com/c-bata/go-prompt"
)

type Mode int

const (
	Keyword Mode = iota
	Node
	Var
)

func (m Mode) Separator() string {
	return []string{" ", ":", "("}[m]
}

type REPL struct {
	prompt      *prompt.Prompt
	complByMode map[Mode]func(prompt.Document) []prompt.Suggest
	mode        Mode
}

func NewREPL() *REPL {
	p := &REPL{}
	p.complByMode = make(map[Mode]func(document prompt.Document) []prompt.Suggest)
	p.prompt = prompt.New(func(in string) {
		fmt.Println("Executor: ", in)
	}, p.completer,
		prompt.OptionSetExitCheckerOnInput(func(in string, breakline bool) bool {
			fmt.Println("ExitChecker ", in, breakline)
			return breakline
		}))
	return p
}

func (r *REPL) Switch(m Mode) func(prompt.Document) []prompt.Suggest {
	r.mode = m
	opt := prompt.OptionCompletionWordSeparator(m.Separator())
	opt(r.prompt)
	if compl, ok := r.complByMode[m]; ok {
		return compl
	}
	return defaultCompl
}

func (r *REPL) Register(m Mode, compl func(prompt.Document) []prompt.Suggest) {
	r.complByMode[m] = compl
}

func (r *REPL) Run() {
	r.prompt.Run()
}

func (r *REPL) Input() string {
	return r.prompt.Input()
}

func (r *REPL) completer(d prompt.Document) []prompt.Suggest {
	line := d.GetWordBeforeCursor()
	compl := defaultCompl
	if strings.HasPrefix(line, ")") || strings.HasPrefix(line, "]") {
		compl = r.Switch(Keyword)
	} else if strings.HasSuffix(line, ":") {
		compl = r.Switch(Node)
	} else if strings.HasSuffix(line, "(") || strings.HasSuffix(line, "[") {
		compl = r.Switch(Var)
	} else {
		compl = r.complByMode[r.mode]
	}
	if compl == nil {
		compl = defaultCompl
	}

	return compl(d)
}

func FindVars(s string) []string {
	subs := cypherVarRegexp.FindAllStringSubmatch(s, -1)
	set := map[string]interface{}{}
	for _, sub := range subs {
		fmt.Println(">>>>" + strings.Join(sub, "^"))
		for i, n := range cypherVarRegexp.SubexpNames() {
			if strings.HasPrefix(n, "var") && sub[i] != "" {
				if sub[i] != "" {
					set[sub[i]] = comp.Present
				}
			}
		}
	}

	var vs = []string{}
	for v := range set {
		vs = append(vs, v)
	}
	sort.Strings(vs)
	return vs
}

func defaultCompl(_ prompt.Document) []prompt.Suggest {
	return nil
}

var cypherVarRegexp *regexp.Regexp

func init() {
	as := `\s+AS\s+(?P<var3>[^\s,]+)`
	rel := `\[(?P<var1>[^]:]+)(:[^]]+)?\]`
	node := `\((?P<var2>[^):]+)(:[^)]+)?\)`
	anonNode := `\((?P<var2>[^):]+)?(:[^)]+)?\)`
	pattern := fmt.Sprintf("(-%s-[<>]?%s)|([\\s-]%s)|(%s)", rel, anonNode, node, as)
	cypherVarRegexp = regexp.MustCompile(pattern)
}
