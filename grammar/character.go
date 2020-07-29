package grammar

import . "github.com/axh432/gogex"

var (
	Underscore       = SetOfCharacters("_")
	Comma            = SetOfCharacters(",")
	OpenBracket      = SetOfCharacters("(")
	ClosedBracket    = SetOfCharacters(")")
	OpenCurlyBrace   = SetOfCharacters("{")
	ClosedCurlyBrace = SetOfCharacters("}")
	Quote            = SetOfCharacters(`"`)
	Dot              = SetOfCharacters(".")
	BackTick         = SetOfCharacters("`")
)
