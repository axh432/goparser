package grammar

import (
	"errors"
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

/*
Loop until you get to VarBlockAssignmentStatement

-> VariableNameList -> VariableName: get value

-> AnyValue -> FunctionCall -> FunctionName: get value
		   -> FunctionCall -> FunctionArgument: get label figure out what type it is

FunctionArgument: get value
FunctionArgument is function call: repeat the function call steps

Sequence: is a struct of the variables in order
Set: is a struct with one of the variables populated
Range: is a slice
Label: is the name and type given to the expression it surrounds

need a symbol table for variable/function names
*/

func ParseVariableName(mt *MatchTree) MatchTree {
	return MatchTree{
		Type:  "Label",
		Label: "VariableName",
		Value: mt.Value,
	}
}

func ParseString(mt *MatchTree) MatchTree {
	return MatchTree{
		Type:     "Label",
		Label:    "String",
		Value:    mt.Value,
		Children: nil,
	}
}

func ParseFunctionArgument(mt *MatchTree) MatchTree {
	child := &mt.Children[0]
	switch child.Label {
	case "VariableName":
		return ParseVariableName(child)
	case "FunctionCall":
		return ParseFunctionCall(child)
	case "String:string":
		return ParseString(child)
	}
	return MatchTree{}
}

func ParseFunctionArguments(mt *MatchTree) []MatchTree {
	var arguments []MatchTree
	for _, child := range mt.Children {
		arguments = append(arguments, ParseFunctionArgument(&child))
	}
	return arguments
}

func ParseFunctionName(mt *MatchTree) MatchTree {
	return MatchTree{
		Type:     "Label",
		Label:    "FunctionName",
		Value:    mt.Value,
		Children: nil,
	}
}

func ParseFunctionCall(mt *MatchTree) MatchTree {
	exp := MatchTree{
		Type:  "Label",
		Label: "FunctionCall",
	}
	for _, child := range mt.Children {
		switch child.Label {
		case "FunctionName":
			exp.Children = append(exp.Children, ParseFunctionName(&child))
		case "FunctionArguments":
			exp.Children = append(exp.Children, ParseFunctionArguments(&child)...)
		}
	}
	return exp
}

func ParseVarBlockAssignmentStatement(mt *MatchTree) MatchTree {
	exp := MatchTree{
		Type:  "Label",
		Label: "VarBlockAssignmentStatement",
	}
	for _, child := range mt.Children {
		switch child.Label {
		case "VariableNameList":
			exp.Children = append(exp.Children, ParseVariableName(&child))
		case "AnyValue":
			exp.Children = append(exp.Children, ParseFunctionArgument(&child))
		}
	}
	return exp
}

func StatementToParser(mt *MatchTree) (error, string) {

	if mt.Label != "VarBlockAssignmentStatement" {
		return errors.New("This function only works for: VarBlockAssignmentStatement"), ""
	}

	variableName := mt.Children[0]

	if variableName.Label != "VariableName" {
		return errors.New("The first child of the statement should be: VariableName"), ""
	}

	switch mt.Children[1].Label {
	case "FunctionCall":
		HandleFunctionCall(&mt.Children[1])
	case "VariableName":
		HandleVariableName(&mt.Children[1])
	}

	return nil, ""
}

func HandleVariableName(mt *MatchTree) {}

func HandleFunctionCall(mt *MatchTree) {
	functionName := mt.Children[0]
	switch functionName.Value {
	case "Label":
		HandleLabel(mt)
	case "SequenceOfCharacters":
		HandleSequenceOfCharacters(mt)
	}

}

func HandleLabel(mt *MatchTree) {
	functionName := mt.Children[0]
	funcCallOrVar := mt.Children[1]
	labelValue := mt.Children[2]
	switch funcCallOrVar.Label {
	case "FunctionCall":
		HandleFunctionCall(&funcCallOrVar)
	case "VariableName":
		HandleVariableName(&funcCallOrVar)
	}
}

func HandleSequenceOfCharacters(mt *MatchTree) {
	functionName := mt.Children[0]
	value := mt.Children[1]
}

func Test_Golang(t *testing.T) {
	t.Run("Golang", func(t *testing.T) {
		var statements = map[string]MatchTree{}
		var statement = MatchTree{}
		fileAsBytes, err := ioutil.ReadFile("./keywords.go")
		require.NoError(t, err)

		tree := Match(string(fileAsBytes), Golang)
		tree = tree.PruneToLabels()

		//fmt.Println(tree.ToMermaidDiagram())

		visit := func(mt *MatchTree) {
			switch mt.Label {
			case "VarBlockAssignmentStatement":
				statement = ParseVarBlockAssignmentStatement(mt)
				statements[statement.Children[0].Value] = statement
			}
		}

		tree.AcceptVisitor(visit)

		//for key, value := range statements {
		//	fmt.Printf("%s: %v+\n", key, value)
		//}

		fmt.Println(statement.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})
}
