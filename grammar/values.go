package grammar

import . "github.com/axh432/gogex"

var (
	DoubleQuotedString = Label(Sequence(Quote, Range(Set(SetOfNotCharacters(`"`), SequenceOfCharacters(`\"`)), 0, -1), Quote), "String:string")
	BackTickString     = Label(Sequence(BackTick, Range(SetOfNotCharacters("`"), 0, -1), BackTick), "String:string")
	StringValue        = Label(Set(DoubleQuotedString, BackTickString), "String:string")
	BoolValue          = Label(Set(SequenceOfCharacters("true"), SequenceOfCharacters("false")), "Bool:bool")
	IntValue           = Label(Range(Number, 1, -1), "Integer:int")
	AnyValue           = Label(Set(StringValue, BoolValue, IntValue, VariableName, FunctionCall), "AnyValue")
)
