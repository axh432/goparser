package tree2struct

import . "github.com/axh432/gogex"

var (
	OpenBracket   = SetOfCharacters("(")
	ClosedBracket = SetOfCharacters(")")
	Comma         = SetOfCharacters(",")
	Hyphen        = SetOfCharacters("-")
	Underscore    = SetOfCharacters("_")
	Equals        = SetOfCharacters("=")

	Int            = Sequence(Range(Hyphen, 0, 1), Range(Number, 1, -1))
	Name           = Range(Set(Letter, Number, Underscore), 1, -1)
	ExpressionName = Label(Name, "name:ExpressionName", "type:string")

	SequenceKeyword             = Label(SequenceOfCharacters("Sequence"), "name:SequenceKeyword", "type:string")
	LabelKeyword                = Label(SequenceOfCharacters("Label"), "name:LabelKeyword", "type:string")
	SetKeyword                  = Label(SequenceOfCharacters("Set"), "name:SetKeyword", "type:string")
	SetOfCharactersKeyword      = Label(SequenceOfCharacters("SetOfCharacters"), "name:SetOfCharactersKeyword", "type:string")
	SetOfNotCharactersKeyword   = Label(SequenceOfCharacters("SetOfNotCharacters"), "name:SetOfNotCharactersKeyword", "type:string")
	SequenceOfCharactersKeyword = Label(SequenceOfCharacters("SequenceOfCharacters"), "name:SequenceOfCharactersKeyword", "type:string")
	RangeKeyword                = Label(SequenceOfCharacters("Range"), "name:RangeKeyword", "type:string")
	VarKeyword                  = Label(SequenceOfCharacters("var"), "name:VarKeyword", "type:string")
	PackageKeyword              = Label(SequenceOfCharacters("package"), "name:PackageKeyword", "type:string")

	SequenceOfCharactersCall = Label(Sequence(SequenceOfCharactersKeyword, OptionalWhitespaceBlock, OpenBracket, OptionalWhitespaceBlock, Label(StringValue, "name:Argument", "type:string"), OptionalWhitespaceBlock, ClosedBracket), "name:SequenceOfCharactersCall", "type:struct")
	SetOfCharactersCall      = Label(Sequence(SetOfCharactersKeyword, OptionalWhitespaceBlock, OpenBracket, OptionalWhitespaceBlock, Label(StringValue, "name:Argument", "type:string"), OptionalWhitespaceBlock, ClosedBracket), "name:SetOfCharactersCall", "type:struct")
	SetOfNotCharactersCall   = Label(Sequence(SetOfNotCharactersKeyword, OptionalWhitespaceBlock, OpenBracket, OptionalWhitespaceBlock, Label(StringValue, "name:Argument", "type:string"), OptionalWhitespaceBlock, ClosedBracket), "name:SetOfNotCharactersCall", "type:struct")

	ExpressionStatement = Label(Sequence(ExpressionName, OptionalWhitespaceBlock, Equals, OptionalWhitespaceBlock, ExpressionArgument), "name:ExpressionStatement", "type:struct")

	VarStatement     = Label(Sequence(VarKeyword, OptionalWhitespaceBlock, Label(ConstructBoundedList(OpenBracket, ExpressionStatement, ClosedBracket), "name:Expressions", "type:[]struct")), "name:VarStatement", "type:struct")
	PackageStatement = Label(Sequence(PackageKeyword, OptionalWhitespaceBlock, Label(Name, "name:PackageName", "type:string")), "name:PackageStatement", "type:struct")

	GogexFile = Label(Sequence(PackageStatement, OptionalWhitespaceBlock, ImportStatement, OptionalWhitespaceBlock, VarStatement), "name:GogexFile", "type:struct")
)

func ExpressionArgument(iter *Iterator) MatchTree {
	return Label(Set(ExpressionName, SequenceCall, SequenceOfCharactersCall, SetOfCharactersCall, SetOfNotCharactersCall, SetCall, RangeCall, LabelCall), "name:ExpressionArgument", "type:struct")(iter)
}

func LabelCall(iter *Iterator) MatchTree {
	labels := Label(ConstructDelimitedList(Label(StringValue, "name:Label", "type:string"), Comma), "name:Labels", "type:[]string")
	return Label(Sequence(LabelKeyword, OptionalWhitespaceBlock, OpenBracket, OptionalWhitespaceBlock, ExpressionArgument, OptionalWhitespaceBlock, Comma, OptionalWhitespaceBlock, labels, OptionalWhitespaceBlock, ClosedBracket), "name:LabelCall", "type:struct")(iter)
}

func RangeCall(iter *Iterator) MatchTree {
	minInt := Label(Int, "name:MinValue", "type:int")
	maxInt := Label(Int, "name:MaxValue", "type:int")
	return Label(Sequence(RangeKeyword, OptionalWhitespaceBlock, OpenBracket, OptionalWhitespaceBlock, ExpressionArgument, OptionalWhitespaceBlock, Comma, OptionalWhitespaceBlock, minInt, OptionalWhitespaceBlock, Comma, OptionalWhitespaceBlock, maxInt, OptionalWhitespaceBlock, ClosedBracket), "name:RangeCall", "type:struct")(iter)
}

func SetCall(iter *Iterator) MatchTree {
	args := Label(ConstructDelimitedBoundedList(OpenBracket, ExpressionArgument, Comma, ClosedBracket), "name:ExpressionArguments", "type:[]struct")
	return Label(Sequence(SetKeyword, OptionalWhitespaceBlock, args), "name:SetCall", "type:struct")(iter)
}

func SequenceCall(iter *Iterator) MatchTree {
	args := Label(ConstructDelimitedBoundedList(OpenBracket, ExpressionArgument, Comma, ClosedBracket), "name:ExpressionArguments", "type:[]struct")
	return Label(Sequence(SequenceKeyword, OptionalWhitespaceBlock, args), "name:SequenceCall", "type:struct")(iter)
}

func ConstructDelimitedBoundedList(openBound Expression, item Expression, delimiter Expression, closedBound Expression) Expression {
	list := Sequence(item, Range(Sequence(OptionalWhitespaceBlock, delimiter, OptionalWhitespaceBlock, item), 0, -1))
	listBounded := Sequence(openBound, OptionalWhitespaceBlock, list, OptionalWhitespaceBlock, closedBound)
	listBoundedEmpty := Sequence(openBound, OptionalWhitespaceBlock, closedBound)
	return Set(listBounded, listBoundedEmpty)
}

func ConstructDelimitedList(item Expression, delimiter Expression) Expression {
	return Sequence(item, Range(Sequence(OptionalWhitespaceBlock, delimiter, OptionalWhitespaceBlock, item), 0, -1))
}

func ConstructBoundedList(openBound Expression, item Expression, closedBound Expression) Expression {
	list := Sequence(item, Range(Sequence(OptionalWhitespaceBlock, item), 0, -1))
	listBounded := Sequence(openBound, OptionalWhitespaceBlock, list, OptionalWhitespaceBlock, closedBound)
	listBoundedEmpty := Sequence(openBound, OptionalWhitespaceBlock, closedBound)
	return Set(listBounded, listBoundedEmpty)
}
