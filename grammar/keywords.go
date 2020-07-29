package grammar

import . "github.com/axh432/gogex"

var (
	StringKeyword   = Label(SequenceOfCharacters("string"), "StringKeyword:string")
	BoolKeyword     = Label(SequenceOfCharacters("bool"), "BoolKeyword:string")
	IntKeyword      = Label(SequenceOfCharacters("int"), "IntKeyword:string")
	VarKeyword      = Label(SequenceOfCharacters("var"), "VarKeyword:string")
	FunctionKeyword = Label(SequenceOfCharacters("func"), "FunctionKeyword:string")
	ImportKeyword   = Label(SequenceOfCharacters("import"), "ImportKeyword:string")
	PackageKeyword  = Label(SequenceOfCharacters("package"), "PackageKeyword:string")
)
