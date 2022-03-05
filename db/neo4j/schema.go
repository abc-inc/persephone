package neo4j

type Cmd struct {
	Name    string
	Desc    string
	SubCmds []Cmd
}

func (c Cmd) String() string {
	return c.Name
}

type Func struct {
	Name     string
	Sig      string
	RetItems []Func
}

func (f Func) String() string {
	return f.Name
}

type Schema struct {
	Labels   []string
	RelTypes []string
	PropKeys []string
	Funcs    []interface{}
	Procs    []Func
	ConCmds  []Cmd
	Params   []string
}

// TODO see neo4jSchema
func NewSchema() *Schema {
	return &Schema{
		Labels:   []string{":State", ":Party", ":Body"},
		RelTypes: []string{":REPRESENTS", ":IS_MEMBER_OF", ":ELECTED_TO"},
		PropKeys: []string{"code", "name", "type"},
		Funcs:    []interface{}{Func{"apoc.coll.avg", "(numbers :: LIST? OF NUMBER?) :: (FLOAT?)", nil}, Func{"apoc.coll.contains", "", nil}},
		Procs:    []Func{{"apoc.algo.aStar", "", nil}},
		ConCmds:  []Cmd{{Name: ":clear"}, {Name: ":play"}, {Name: ":help"}},
		Params:   []string{"age", "name", "surname"},
	}
}
