package grammar

import . "github.com/axh432/gogex"

var (
	ImportAccessOperator  = Label(Range(SetOfCharacters("_."), 0, 1), "ImportAccessOperator:string")
	AssignmentOperator    = Label(SequenceOfCharacters(":="), "AssignmentOperator")
	VarAssignmentOperator = Label(SequenceOfCharacters("="), "VarAssignmentOperator")
	AccessOperator        = Label(Dot, "AccessOperator")
)
