package main

import (
	"github.com/abc-inc/merovingian/repl"
)

func main() {
	//q := heredoc.Doc(`
//PROFILE  MATCH (n:Node)--(o)
//WITH n, count(o) AS cnt
//RETURN n.name, cnt`)
	repl.NewRepl()
}
