package console

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers"
)

// CypherLexer highlights syntax tokens using Chroma.
var CypherLexer = lexers.Register(MustNewLazyLexer(
	&Config{
		Name:            "Cypher",
		Aliases:         []string{"cypher"},
		Filenames:       []string{"*.cyp", "*.cypher"},
		MimeTypes:       []string{},
		CaseInsensitive: true,
	},
	func() Rules {
		return Rules{
			"root": {
				Include("comment"),
				Include("keywords"),
				Include("clauses"),
				Include("relations"),
				Include("strings"),
				Include("whitespace"),
				Include("barewords"),
			},
			"comment": {
				{Pattern: `^.*//.*\n`, Type: CommentSingle},
			},
			"keywords": {
				{Pattern: `(create|order|match|limit|set|skip|start|return|with|where|delete|foreach|not|by|true|false)\b`, Type: Keyword},
			},
			"clauses": {
				{Pattern: `(all|any|as|asc|ascending|assert|call|case|create|create\s+index|create\s+unique|delete|desc|descending|distinct|drop\s+constraint\s+on|drop\s+index\s+on|end|ends\s+with|fieldterminator|foreach|in|is\s+node\s+key|is\s+null|is\s+unique|limit|load\s+csv\s+from|match|merge|none|not|null|on\s+match|on\s+create|optional\s+match|order\s+by|remove|return|set|skip|single|start|starts\s+with|then|union|union\s+all|unwind|using\s+periodic\s+commit|yield|where|when|with)\b`, Type: Keyword},
			},
			"relations": {
				{Pattern: `(-\[)(.*?)(\]->)`, Type: ByGroups(Operator, UsingSelf("root"), Operator)},
				{Pattern: `(<-\[)(.*?)(\]-)`, Type: ByGroups(Operator, UsingSelf("root"), Operator)},
				{Pattern: `(-\[)(.*?)(\]-)`, Type: ByGroups(Operator, UsingSelf("root"), Operator)},
				{Pattern: `-->|<--|\[|\]`, Type: Operator},
				{Pattern: `<|>|<>|=|<=|=>|\(|\)|\||:|,|;`, Type: Punctuation},
				{Pattern: `[.*{}]`, Type: Punctuation},
			},
			"strings": {
				{Pattern: `"(?:\\[tbnrf\'"\\]|[^\\"])*"`, Type: LiteralString},
				{Pattern: "`(?:``|[^`])+`", Type: NameVariable},
			},
			"whitespace": {
				{Pattern: `\s+`, Type: TextWhitespace},
			},
			"barewords": {
				{Pattern: `[a-z]\w*`, Type: Name},
				{Pattern: `\d+`, Type: LiteralNumber},
			},
		}
	},
))
