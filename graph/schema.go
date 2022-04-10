package graph

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
	Funcs    []Func
	Procs    []Func
	ConCmds  []Cmd
	Params   []string
}
