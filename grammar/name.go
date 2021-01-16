package grammar

import . "github.com/axh432/gogex"

var (
	Name         = Range(Set(Letter, Number, Underscore), 1, -1)
	FunctionName = Label(Name, "FunctionName")
	ImportName   = Label(Name, "ImportName:string")
	PackageName  = Label(Name, "PackageName:string")
	ParentName   = Label(Name, "ParentName")
	TypeName     = Label(Name, "TypeName")
	VariableName = Label(Name, "VariableName")
)
