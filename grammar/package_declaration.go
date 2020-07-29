package grammar

import . "github.com/axh432/gogex"

var (
	PackageDeclaration = Label(Sequence(PackageKeyword, OptionalWhitespaceBlock, PackageName), "PackageDeclaration")
)
