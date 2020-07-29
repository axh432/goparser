package grammar

import . "github.com/axh432/gogex"

var (
	VarBlockAssignmentStatement = Label(Sequence(VariableNameList, OptionalWhitespaceBlock, OptionalTypeName, OptionalWhitespaceBlock, VarAssignmentOperator, OptionalWhitespaceBlock, AnyValue), "VarBlockAssignmentStatement")
	VarBlockList                = Label(ConstructBoundedList(OpenBracket, VarBlockAssignmentStatement, ClosedBracket), "VarBlockAssignmentStatements")
	VarBlock                    = Label(Sequence(VarKeyword, OptionalWhitespaceBlock, VarBlockList), "VarBlock")
)
