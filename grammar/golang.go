package grammar

import . "github.com/axh432/gogex"

var (
	TypeOptions = Set(VarDeclaration, FunctionDeclaration)
	Golang      = Label(Sequence(PackageDeclaration, OptionalWhitespaceBlock, ImportDeclaration, OptionalWhitespaceBlock, TypeOptions), "Golang")
)
