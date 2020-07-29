package grammar

import (
	. "github.com/axh432/gogex"
)

var (
	FunctionParameter     = Label(Sequence(VariableName, OptionalWhitespaceBlock, TypeName), "FunctionParameter")
	FunctionParameterList = Label(ConstructDelimitedBoundedList(OpenBracket, FunctionParameter, Comma, ClosedBracket), "FunctionParameters")

	ReturnType               = Label(Name, "ReturnType")
	ReturnParametersNamed    = Label(FunctionParameterList, "ReturnParametersNamed")
	ReturnParametersSingle   = ReturnType
	ReturnParametersList     = ConstructDelimitedBoundedList(OpenBracket, ReturnType, Comma, ClosedBracket)
	ReturnParameters         = Label(Set(ReturnParametersSingle, ReturnParametersList, ReturnParametersNamed), "ReturnParameters")
	OptionalReturnParameters = Range(ReturnParameters, 0, 1)

	FunctionSignature = Label(Sequence(FunctionKeyword, OptionalWhitespaceBlock, FunctionName, OptionalWhitespaceBlock, FunctionParameterList, OptionalWhitespaceBlock, OptionalReturnParameters), "FunctionSignature")
)
