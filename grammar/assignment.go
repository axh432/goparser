package grammar

import . "github.com/axh432/gogex"

var (
	AssignmentStatement = Label(Sequence(VariableNameList, OptionalWhitespaceBlock, AssignmentOperator, OptionalWhitespaceBlock, AnyValue), "AssignmentStatement")
)
