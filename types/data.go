package types

type Data struct {
	Type              Type
	Path              []string
	FilterLastElement bool
}

// AllCompData is the default.
var AllCompData []Data

var AllComp = []Type{Variable, Parameter, PropertyKey, FunctionName, Keyword}

func init() {
	for _, t := range AllComp {
		AllCompData = append(AllCompData, Data{Type: t})
	}
}
