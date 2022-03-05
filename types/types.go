package types

type Type int

const (
	Keyword Type = iota
	Label
	Variable
	Parameter
	PropertyKey
	RelationshipType
	FunctionName
	ProcedureName
	ConsoleCommandName
	ConsoleCommandSubCommand
	ProcedureOutput
	// Noop returns no autocompletion
	Noop
)

// All is the default.
var All = []Type{Variable, Parameter, PropertyKey, FunctionName, Keyword}

func (t Type) String() string {
	return []string{
		"keyword", "label", "variable", "parameter", "propertyKey", "relationshipType",
		"function", "procedure", "consoleCommand", "consoleCommandSubcommand", "procedureOutput",
	}[t]
}
