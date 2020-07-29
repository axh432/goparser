package grammar

import (
	. "github.com/axh432/gogex"
)

func FunctionCall(iter *Iterator) MatchTree {
	FunctionCallArgument := Label(Set(StringValue, BoolValue, IntValue, VariableName, FunctionCall), "FunctionArgument")
	FunctionCallArgumentList := Label(ConstructDelimitedBoundedList(OpenBracket, FunctionCallArgument, Comma, ClosedBracket), "FunctionArguments")
	OptionalParentName := Range(Label(Sequence(PackageName, OptionalWhitespaceBlock, Dot, OptionalWhitespaceBlock), "ParentName"), 0, 1)
	return Label(Sequence(OptionalParentName, FunctionName, OptionalWhitespaceBlock, FunctionCallArgumentList), "FunctionCall")(iter)
}

/*var(
	FunctionCallArgument = Label(AnyValue, "argument")
	FunctionCallArgumentList = ConstructDelimitedBoundedList(OpenBracket, FunctionCallArgument, Comma, ClosedBracket)
	OptionalPackageName = Range(Sequence(PackageName, OptionalWhitespaceBlock, Dot, OptionalWhitespaceBlock), 0, 1)
	FunctionCall = Sequence(OptionalPackageName, FunctionName, OptionalWhitespaceBlock, FunctionCallArgumentList)
)*/
