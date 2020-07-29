package grammar

import . "github.com/axh432/gogex"

var (
	FunctionDeclaration = Label(Sequence(FunctionSignature, OptionalWhitespaceBlock, FunctionBody), "FunctionDeclaration")
)
