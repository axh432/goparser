package grammar

import . "github.com/axh432/gogex"

var (
	ImportStatement   = Label(Sequence(ImportAccessOperator, OptionalWhitespaceBlock, StringValue), "Import")
	ImportList        = Label(ConstructBoundedList(OpenBracket, ImportStatement, ClosedBracket), "Imports:[]Import")
	ImportSingle      = Label(ImportStatement, "Imports:[]Import")
	ImportDeclaration = Label(Sequence(ImportKeyword, OptionalWhitespaceBlock, Set(ImportList, ImportSingle)), "ImportDeclaration")
)
