package grammar

import . "github.com/axh432/gogex"

var (
	OptionalTypeName       = Range(TypeName, 0, 1)
	VarAssignmentStatement = Label(Sequence(
		VarKeyword,
		OptionalWhitespaceBlock,
		VariableNameList,
		OptionalWhitespaceBlock,
		OptionalTypeName,
		OptionalWhitespaceBlock,
		VarAssignmentOperator,
		OptionalWhitespaceBlock,
		AnyValue), "VarAssignmentStatement")
)
