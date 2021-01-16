package tree2struct

import . "github.com/axh432/gogex"

var (
	Quote                   = SetOfCharacters(`"`)
	BackTick                = SetOfCharacters("`")
	BackTickString          = Sequence(BackTick, Range(SetOfNotCharacters("`"), 0, -1), BackTick)
	DoubleQuotedString      = Sequence(Quote, Range(Set(SetOfNotCharacters(`"`), SequenceOfCharacters(`\"`)), 0, -1), Quote)
	StringValue             = Set(DoubleQuotedString, BackTickString)
	ImportKeyword           = Label(SequenceOfCharacters("import"), "name:ImportKeyword", "type:string")
	OptionalWhitespaceBlock = Range(Whitespace, 0, -1)
	ImportAccessOperator    = Label(Range(SetOfCharacters("_."), 0, 1), "name:ImportAccessOperator", "type:string")
	ImportStatement         = Label(Sequence(ImportKeyword, OptionalWhitespaceBlock, ImportAccessOperator, OptionalWhitespaceBlock, Label(StringValue, "name:ImportURL", "type:string")), "name:ImportStatement", "type:struct")
)
