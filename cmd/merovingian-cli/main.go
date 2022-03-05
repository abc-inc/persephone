package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/abc-inc/merovingian/_deprecated"
	"github.com/abc-inc/merovingian/comp"
	"github.com/abc-inc/merovingian/db"
	graph "github.com/abc-inc/merovingian/db/neo4j"
	"github.com/abc-inc/merovingian/parser"
	"github.com/abc-inc/merovingian/types"
	"github.com/abc-inc/merovingian/web"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/c-bata/go-prompt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var present = struct{}{}

type exampleListener struct {
	Vars   map[string]interface{}
	Nodes  []string
	Labels []string
	*parser.BaseCypherListener
	Statements       []*parser.CypherPartContext
	InConsoleCommand bool
	Queries          []*parser.CypherQueryContext
	Indexes          map[types.Type]*comp.Index
}

func (l *exampleListener) PrintChildren(cs []antlr.Tree, n int) {
	for _, c := range cs {
		fmt.Println(strings.Repeat(" ", n), "'"+c.(antlr.ParseTree).GetText()+"'", reflect.TypeOf(c))
		switch c.(type) {
		case *parser.ReturnItemContext:
			ric := c.GetChild(c.GetChildCount() - 1)
			switch ric := ric.(type) {
			case *parser.VariableContext:
				fmt.Println(strings.Repeat("#", 10), "RIC", ric.GetText())
				l.Vars[ric.GetText()] = present
			case *parser.ExpressionContext:
				fmt.Println(strings.Repeat("#", 10), "RIC", ric.GetText())
			default:
				panic(ric)
			}
		case *parser.VariableContext:
			fmt.Println(strings.Repeat("#", 10), "VARIABLE", c.GetChild(0).(*parser.SymbolicNameContext).GetText())
		default:
			l.PrintChildren(c.GetChildren(), n+1)
		}
	}
}

func NewExampleListener() *exampleListener {
	return &exampleListener{Vars: make(map[string]interface{})}
}

func (l *exampleListener) EnterCypherPart(ctx *parser.CypherPartContext) {
	fmt.Println("> PART", ctx.GetText())
}

func (l *exampleListener) ExitCypherPart(ctx *parser.CypherPartContext) {
	if len(l.Statements) == 0 {
		l.Statements = append(l.Statements, ctx)
	}
}

func (l *exampleListener) EnterCypherConsoleCommand(ctx *parser.CypherConsoleCommandContext) {
	for _, v := range l.Indexes {
		v.AddQuery()
	}
	l.InConsoleCommand = true
}

func (l *exampleListener) ExitCypherConsoleCommand() {
	l.InConsoleCommand = false
}

func (l *exampleListener) EnterCypherQuery(ctx *parser.CypherQueryContext) {
	l.Queries = append(l.Queries, ctx)
	for _, v := range l.Indexes {
		v.AddQuery()
	}
}

func (l *exampleListener) EnterVariable(ctx *parser.VariableContext) {
	l.Vars[ctx.GetText()] = present
}

func (l *exampleListener) ExitVariable(ctx *parser.VariableContext) {
	if l.InConsoleCommand {
		return
	}
	l.Indexes[types.Variable].AddVariable(ctx)
}

func (l *exampleListener) EnterPattern(ctx *parser.PatternContext) {
	fmt.Println("> PATTERN", reflect.TypeOf(ctx.GetChild(0).(*parser.PatternPartContext).GetChild(0).(*parser.AnonymousPatternPartContext).GetChild(0).(*parser.PatternElementContext).GetChild(0)))
	l.PrintChildren(ctx.GetChildren(), 0)
}

func (l *exampleListener) EnterWithClause(ctx *parser.WithClauseContext) {
	l.Vars = make(map[string]interface{})
	fmt.Println("> WITH_CLAUSE", ctx.GetText())
	l.PrintChildren(ctx.GetChildren(), 0)
}

func (l *exampleListener) EnterReturnBody(ctx *parser.ReturnBodyContext) {
	l.Vars = make(map[string]interface{})
	fmt.Println("> RETURN", ctx.GetText())
	l.PrintChildren(ctx.GetChildren(), 0)
}

func main() {
	dbUri := "neo4j://localhost:7687"
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth("neo4j", "root", ""))
	if err != nil {
		panic(err)
	}
	conn := graph.NewConn(driver)

	es, err := conn.Metadata()
	if err != nil {
		log.Fatalln(err)
	}

	l := NewExampleListener()
	for _, e := range es {
		l.Nodes = append(l.Nodes, e.Name)
	}

	// Setup the input
	//is := antlr.NewInputStream("MATCH (n:Node) WITH max(n) AS max, count(n) AS count RETURN count, max")
	is := antlr.NewInputStream(`
match (prgs:PROGRAM_STATUS)
with prgs limit 100
with collect(distinct prgs.MBR_DK) as mbr_dks
unwind mbr_dks as mbr_dk
optional match (s:PROGRAM_STATUS {MBR_DK:mbr_dk})-[:UPDATED_TO*]->(e:PROGRAM_STATUS)
where not (()-[]->(s)) and not ((e)-[]->()) 
with mbr_dk, 
	case when s is not null
    	then collect({MBR_DK:s.MBR_DK, CM_CASE_ID:s.CM_CASE_ID, ENROLL_START_DT:s.ENROLL_START_DT, ENROLL_END_DT:coalesce(e.ENROLL_END_DT,date('9999-01-01'))})
    	else [] end 
	as records 
optional match (s:PROGRAM_STATUS {MBR_DK:mbr_dk})
where not (()-[]->(s)) and not ((s)-[]->())
with
	case when s is not null
    	then records + collect({MBR_DK:s.MBR_DK, CM_CASE_ID:s.CM_CASE_ID, ENROLL_START_DT:s.ENROLL_START_DT, ENROLL_END_DT:coalesce(s.ENROLL_END_DT,date('9999-01-01'))})
        else records end
	as allRecords
unwind allRecords as records
with records order by records.ENROLL_START_DT, records.ENROLL_END_DT
with collect(records)[0..count(records)-1] as r_collection_head, tail(collect(records)) as r_collection_tail
unwind r_collection_head as r1
with r1, r_collection_tail, apoc.coll.indexOf(r_collection_head, r1) as i
with r1, r_collection_tail[i] as r2
with r1,r2, max(r1.ENROLL_START_DT, r2.ENROLL_START_DT) < min(r1.ENROLL_END_DT, r2.ENROLL_END_DT) as has_overlap
with collect ([{
			MBR_DK:r1.MBR_DK
            ,CM_CASE_ID:r1.CM_CASE_ID + ':' + r2.CM_CASE_ID
			,R1_ENROLL_START_DT:r1.ENROLL_START_DT
            ,R1_ENROLL_END_DT:r1.ENROLL_END_DT
            ,R2_ENROLL_START_DT:r2.ENROLL_START_DT
            ,R2_ENROLL_END_DT:r2.ENROLL_END_DT
            ,HAS_OVERLAP:has_overlap}]) as results
return results
`)

	// Create the Lexer
	lexer := parser.NewCypherLexer(is)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the JSON Parser
	par := parser.NewCypherParser(stream)

	// Finally, walk the tree
	v := par.CypherQuery()
	antlr.ParseTreeWalkerDefault.Walk(l, v)
	//fmt.Println(v.GetChildren()[0].(*parser.RegularQueryContext).GetChild(0).(*parser.SingleQueryContext).GetChild(0).(*parser.ClauseContext).GetChild(0).(*parser.MatchClauseContext).GetChild(2).(*parser.PatternContext).GetText())
	fmt.Println(l.Vars)
	fmt.Println(l.Nodes)

	if true {
		return
	}

	// Read all tokens
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}

	ls := []prompt.Suggest{}
	for _, e := range es {
		ls = append(ls, prompt.Suggest{Text: e.Name})
	} // https://godoc.org/bramp.net/antlr4/json#BaseJSONListener

	p := _deprecated.NewREPL()
	p.Register(_deprecated.Node, func(d prompt.Document) []prompt.Suggest {
		sub := d.GetWordBeforeCursor()
		sub = sub[strings.IndexByte(sub, ':')+1:]
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>" + sub)
		return prompt.FilterHasPrefix(ls, sub, false)
	})

	p.Register(_deprecated.Var, func(d prompt.Document) []prompt.Suggest {
		vars := _deprecated.FindVars(d.TextBeforeCursor())
		var ss []prompt.Suggest
		for _, v := range vars {
			ss = append(ss, prompt.Suggest{Text: v})
		}
		return ss
	})

	cyp := p.Input()
	fmt.Println("CYP: " + cyp)

	err = web.Foo(os.Stdout, conn, db.Request{Query: cyp})
	fmt.Println(err)
}
