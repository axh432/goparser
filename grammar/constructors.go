package grammar

import . "github.com/axh432/gogex"

func ConstructList(item Expression) Expression {
	return Sequence(item, Range(Sequence(OptionalWhitespaceBlock, item), 0, -1))
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

func ConstructDelimitedBoundedList(openBound Expression, item Expression, delimiter Expression, closedBound Expression) Expression {
	list := Sequence(item, Range(Sequence(OptionalWhitespaceBlock, delimiter, OptionalWhitespaceBlock, item), 0, -1))
	listBounded := Sequence(openBound, OptionalWhitespaceBlock, list, OptionalWhitespaceBlock, closedBound)
	listBoundedEmpty := Sequence(openBound, OptionalWhitespaceBlock, closedBound)
	return Set(listBounded, listBoundedEmpty)
}
