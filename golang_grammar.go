package goparser

import (
	_ "fmt"
	. "github.com/axh432/gogex"
	_ "strings"
)

var (

	//Primitives
	whitespaceBlock                  = Range(Whitespace, 1, -1)
	optionalWhitespaceBlock          = Range(Whitespace, 0, -1)
	whitespaceNoNewLine              = SetOfCharacters(" \t")
	whitespaceNoNewLineBlock         = Range(whitespaceNoNewLine, 1, -1)
	optionalWhitespaceNoNewLineBlock = Range(SetOfCharacters(" \t"), 0, -1)
	newline                          = SequenceOfCharacters("\n")
	whitespaceAtLeastOneNewLineBlock = Sequence(optionalWhitespaceNoNewLineBlock, newline, optionalWhitespaceBlock)
	underscore                       = SetOfCharacters("_")
	comma                            = SetOfCharacters(",")
	openBracket                      = SetOfCharacters("(")
	closedBracket                    = SetOfCharacters(")")
	openCurlyBrace                   = SetOfCharacters("{")
	closedCurlyBrace                 = SetOfCharacters("}")
	quote                            = SetOfCharacters(`"`)
	dot                              = SetOfCharacters(".")
	word                             = Range(Letter, 1, -1)
	String                           = Sequence(quote, Range(Set(SetOfNotCharacters(`"`), SequenceOfCharacters(`\"`)), 1, -1), quote)
	boolValue                        = Set(SequenceOfCharacters("true"), SequenceOfCharacters("false"))
	integerValue                     = Range(Number, 1, -1)
	listOfIntegerValues              = Sequence(Range(Sequence(integerValue, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), integerValue)

	//name
	letterNumberUnderscoreBlock         = Range(Set(Letter, Number, underscore), 1, -1)
	optionalLetterNumberUnderscoreBlock = Range(Set(Letter, Number, underscore), 0, -1)
	letterOrUnderscore                  = Set(Letter, underscore)
	name                                = Sequence(letterOrUnderscore, optionalLetterNumberUnderscoreBlock)

	typeName     = Label(name, "typename")
	variableName = Label(name, "variablename")
	returnType   = Label(name, "returntype")
	functionName = Label(name, "functionName")

	//Import
	Import                  = SequenceOfCharacters("import")
	importNameWithSpecifier = Sequence(SetOfCharacters("_."), optionalWhitespaceNoNewLineBlock, String)
	importNameNoSpecifier   = String
	importName              = Set(importNameWithSpecifier, importNameNoSpecifier)
	importMultiple          = Sequence(Range(Sequence(importName, whitespaceAtLeastOneNewLineBlock), 1, -1), importName)
	importBoundedMultiple   = Sequence(openBracket, optionalWhitespaceBlock, importMultiple, optionalWhitespaceBlock, closedBracket)
	importBoundedSingle     = Sequence(openBracket, optionalWhitespaceBlock, importName, optionalWhitespaceBlock, closedBracket)
	importBoundedEmpty      = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	importBoundedAll        = Set(importBoundedMultiple, importBoundedSingle, importBoundedEmpty)
	importSingle            = importName
	importDeclaration       = Sequence(Import, optionalWhitespaceBlock, Set(importBoundedAll, importSingle))

	//Function Signature
	Func                          = SequenceOfCharacters("func")
	parameter                     = Label(Sequence(variableName, whitespaceNoNewLineBlock, typeName), "parameter")
	functionParametersList        = Sequence(Range(Sequence(parameter, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), optionalWhitespaceBlock, parameter)
	functionParametersBoundedList = Sequence(openBracket, optionalWhitespaceBlock, functionParametersList, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionParametersSingle      = Sequence(openBracket, optionalWhitespaceBlock, parameter, optionalWhitespaceBlock, closedBracket)
	functionParametersEmpty       = Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	functionParameters            = Set(functionParametersBoundedList, functionParametersSingle, functionParametersEmpty)

	//Function Return Parameters
	returnParametersNamed       = functionParameters
	returnParametersSingle      = returnType
	returnParametersList        = Sequence(Range(Sequence(returnType, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), optionalWhitespaceBlock, returnType)
	returnParametersBoundedList = Sequence(openBracket, optionalWhitespaceBlock, returnParametersList, optionalWhitespaceNoNewLineBlock, closedBracket)
	returnParameters            = Set(returnParametersSingle, returnParametersBoundedList, returnParametersNamed)
	optionalReturnParameters    = Range(returnParameters, 0, 1)

	//Function Signature
	functionSignature = Sequence(Func, whitespaceNoNewLineBlock, functionName, optionalWhitespaceNoNewLineBlock, functionParameters, optionalWhitespaceNoNewLineBlock, optionalReturnParameters)

	//Var Assign Statement
	Var                   = SequenceOfCharacters("var")
	varAssignmentOperator = SetOfCharacters("=")
	valuePossibilities    = Set(String, boolValue, integerValue, listOfIntegerValues, functionCall)
	optionalTypeName      = Range(typeName, 0, 1)
	varNameList           = Sequence(Range(Sequence(variableName, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), variableName)
	varNames              = Set(varNameList, variableName)
	varAssignStatement    = Sequence(Var, optionalWhitespaceBlock, varNames, whitespaceNoNewLineBlock, optionalTypeName, optionalWhitespaceNoNewLineBlock, varAssignmentOperator, optionalWhitespaceBlock, valuePossibilities)

	//Var Declaration Statement
	varDeclarationStatement = Sequence(Var, whitespaceBlock, varNames, whitespaceNoNewLineBlock, typeName)

	//Var Statement
	varStatement = Set(varAssignStatement, varDeclarationStatement)

	//Assign Statement
	assignmentOperator = SequenceOfCharacters(":=")
	assignStatement    = Sequence(varNames, optionalWhitespaceNoNewLineBlock, assignmentOperator, optionalWhitespaceBlock, valuePossibilities)

	//Function Body
	statement              = Set(varStatement, assignStatement, functionCall)
	statements             = Label(Sequence(statement, Range(Sequence(whitespaceAtLeastOneNewLineBlock, statement), 0, -1)), "Statements")
	statementsBounded      = Sequence(openCurlyBrace, optionalWhitespaceBlock, statements, optionalWhitespaceBlock, closedCurlyBrace)
	statementsBoundedEmpty = Sequence(openCurlyBrace, optionalWhitespaceBlock, closedCurlyBrace)
	functionBody           = Set(statementsBounded, statementsBoundedEmpty)

	//Function
	functionDeclaration = Sequence(functionSignature, optionalWhitespaceNoNewLineBlock, functionBody)

	//Package
	Package            = SequenceOfCharacters("package")
	packageName        = Label(name, "packagename")
	packageDeclaration = Sequence(Package, whitespaceNoNewLineBlock, packageName)

	//Basic Golang
	basicGo = Sequence(packageDeclaration, whitespaceAtLeastOneNewLineBlock, importDeclaration, whitespaceAtLeastOneNewLineBlock, functionDeclaration)
)

//Function Call
func functionCall(iter *Iterator) MatchTree {
	functionCallParameter := Set(variableName, String, functionCall)
	functionCallParametersMultiple := Sequence(Range(Sequence(functionCallParameter, optionalWhitespaceNoNewLineBlock, comma, optionalWhitespaceBlock), 1, -1), functionCallParameter)
	functionCallParametersBoundedMultiple := Sequence(openBracket, optionalWhitespaceBlock, functionCallParametersMultiple, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionCallParametersBoundedSingle := Sequence(openBracket, optionalWhitespaceBlock, functionCallParameter, optionalWhitespaceNoNewLineBlock, closedBracket)
	functionCallParametersBoundedEmpty := Sequence(openBracket, optionalWhitespaceBlock, closedBracket)
	functionCallParametersBoundedAll := Set(functionCallParametersBoundedMultiple, functionCallParametersBoundedSingle, functionCallParametersBoundedEmpty)
	optionalPackageName := Range(Sequence(packageName, optionalWhitespaceNoNewLineBlock, dot, optionalWhitespaceBlock), 0, 1)

	return Sequence(optionalPackageName, functionName, optionalWhitespaceNoNewLineBlock, functionCallParametersBoundedAll)(iter)
}
