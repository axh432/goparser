package grammar

import . "github.com/axh432/gogex"

var (
	VariableNameList = Label(ConstructDelimitedList(VariableName, Comma), "VariableNameList")
	VarStatement     = Label(Sequence(VarKeyword, OptionalWhitespaceBlock, VariableNameList, OptionalWhitespaceBlock, TypeName), "VarStatement")
	VarDeclaration   = Label(Set(VarAssignmentStatement, VarStatement, VarBlock), "VarDeclaration:struct")
)
