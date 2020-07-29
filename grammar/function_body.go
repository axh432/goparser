package grammar

import . "github.com/axh432/gogex"

var (
	Statement    = Label(Set(VarStatement, AssignmentStatement, FunctionCall), "Statement")
	FunctionBody = Label(ConstructBoundedList(OpenCurlyBrace, Statement, ClosedCurlyBrace), "FunctionBody")
)
