package grammar

import (
	"errors"
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/axh432/goparser/grammar/modelvisitor"
	"io/ioutil"
	"strconv"
	"strings"
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
SequenceOfCharacters: is a string
SetOfCharacters: is a string
Label: is the name and type given to the expression it surrounds

need a symbol table for variable/function names
*/

var statements = map[string]MatchTree{}
var typeDefinitions = map[string]modelvisitor.TypeDefinition{}

//------------------------------------------------

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

//--------------------------------------------------

//if function call is a label -> get the name and determine what it wraps around
//if function call is a sequence -> then we are a struct
//for each child - find their type
//if function call is a set -> then we are a struct
//for each child - find their type
//if function call is a range -> then we are an array
//for each child - find their type
//if function call is a sequence of characters -> then this is a string
//if function call is a SetOfCharacters -> then this is a string
//if variable name -> then

const STRING, ARRAY, STRUCT int = 0, 1, 2

func createDefinitionFromStatement(statement *MatchTree) modelvisitor.TypeDefinition {
	expression := getExpressionFromStatement(statement)
	name := getNameFromStatement(statement)
	if isALabel(expression) {
		return createDefinitionFromExpression(getSubExpressionFromLabel(expression), getNameFromLabel(expression))
	}
	return createDefinitionFromExpression(expression, name)
}

func createDefinitionFromExpression(wrappedFunc *MatchTree, typeName string) modelvisitor.TypeDefinition {
	dataType := determineDataType(wrappedFunc)
	switch dataType {
	case STRING:
		return modelvisitor.PrimitiveDefinition{
			Name: typeName,
			Type: "string",
		}
	case ARRAY:
		return modelvisitor.ArrayDefinition{
			Name: typeName,
			Type: getArrayType(wrappedFunc),
		}
	case STRUCT:
		return modelvisitor.StructDefinition{
			Name:      typeName,
			Type:      typeName,
			Variables: getStructVariables(wrappedFunc),
		}
	}
	return nil
}

func getStructVariables(wrappedFunc *MatchTree) map[string]string {
	variables := map[string]string{}
	numberOfChildren := len(wrappedFunc.Children) - 1

	for i := 1; i < numberOfChildren; i++ {
		child := &wrappedFunc.Children[i]
		switch child.Label {
		case "FunctionCall":
			if isALabel(child) {
				name := getNameFromLabel(child)
				typeDef := createDefinitionFromExpression(getSubExpressionFromLabel(child), name)
				typeDefinitions[name] = typeDef
				variables[name] = typeDef.GetType()
			}
			panic("encountered a function call that was not wrapped in a label: " + child.Children[0].Value)
		case "VariableName":
			typeDef := findDefinition(child.Value)
			variables[child.Value] = typeDef.GetType()
		}
	}

	return variables
}

func getArrayType(wrappedFunc *MatchTree) string {
	expression := &wrappedFunc.Children[1]
	switch expression.Label {
	case "FunctionCall":
		if isALabel(expression) {
			name := getNameFromLabel(expression)
			typeDef := createDefinitionFromExpression(getSubExpressionFromLabel(expression), name)
			typeDefinitions[name] = typeDef
			return fmt.Sprintf("[]%s", typeDef.GetType())
		}
		panic("encountered a function call that was not wrapped in a label: " + expression.Children[0].Value)
	case "VariableName":
		typeDef := findDefinition(expression.Value)
		return fmt.Sprintf("[]%s", typeDef.GetType())
	}
	panic("unable to find array type, label is: " + wrappedFunc.Children[1].Label)
}

func findDefinition(name string) modelvisitor.TypeDefinition {
	if typeDef, ok := typeDefinitions[name]; ok {
		return typeDef
	}

	if statement, ok := statements[name]; ok {
		typeDef := createDefinitionFromStatement(&statement)
		typeDefinitions[name] = typeDef
		return typeDef
	}

	panic("can't find the type definition or the statement for: " + name)
}

func determineDataType(wrappedFunc *MatchTree) int {
	switch wrappedFunc.Children[0].Value {
	case "Sequence", "set":
		return STRUCT
	case "Range":
		return ARRAY
	default:
		return STRING
	}
}

func getNameFromStatement(statement *MatchTree) string {
	return statement.Children[0].Value
}

func getExpressionFromStatement(statement *MatchTree) *MatchTree {
	return &statement.Children[1]
}

func getSubExpressionFromLabel(labelNode *MatchTree) *MatchTree {
	return &labelNode.Children[1]
}

func getNameFromLabel(labelNode *MatchTree) string {
	return labelNode.Children[2].Value
}

func isALabel(expression *MatchTree) bool {
	return expression.Label == "FunctionCall" &&
		expression.Children[0].Value == "Label"
}

//----------------------------
// matchtree path

func matchTreePath(mt *MatchTree, path string) (*MatchTree, error) {
	return matchTreePathRecursive(mt, strings.Split(path, "/"))
}

func matchTreePathRecursive(mt *MatchTree, elements []string) (*MatchTree, error) {
	indexOfLastElement := len(mt.Children) - 1

	i, err := strconv.Atoi(elements[0])
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Match Tree Path should contain only numbers. Found: %s", elements[0]))
	}

	if i > indexOfLastElement {
		return nil, errors.New(fmt.Sprintf("Attempted to find a child at index: %d, when the highest element index is: %d", i, indexOfLastElement))
	}

	if len(elements) == 1 {
		return &mt.Children[i], nil
	}

	return matchTreePathRecursive(&mt.Children[i], elements[1:indexOfLastElement])
}

//-----------------------------
func HandleStatement(mt *MatchTree) (error, *modelvisitor.TypeDefinition) {

	if mt.Label != "VarBlockAssignmentStatement" {
		return errors.New("This function only works for: VarBlockAssignmentStatement"), nil
	}

	variableName := mt.Children[0]

	if variableName.Label != "VariableName" {
		return errors.New("The first child of the statement should be: VariableName"), nil
	}

	switch mt.Children[1].Label {
	case "FunctionCall":
		HandleFunctionCall(&mt.Children[1])
	case "VariableName":
		HandleVariableName(&mt.Children[1])
	}

	return nil, nil
}

func HandleVariableName(mt *MatchTree) {
	if variable, ok := statements[mt.Value]; ok {
		HandleStatement(&variable)
	} else {
		panic(fmt.Sprintf("unknown statement '%s' you need to read it in first", mt.Value))
	}

}

func HandleFunctionCall(mt *MatchTree) {
	functionName := mt.Children[0]
	switch functionName.Value {
	case "Label":
		HandleLabel(mt)
	case "SequenceOfCharacters":
		HandleSequenceOfCharacters(mt)
	case "Sequence":
		HandleSequence(mt)
	}

}

func HandleLabel(mt *MatchTree) {
	functionName := mt.Children[0]
	funcCallOrVar := mt.Children[1]
	labelValue := mt.Children[2]

	if functionName.Value != "Label" {
		//throw a fit
		return
	}

	if strings.Contains(labelValue.Value, ":string") {
		return
	}

	switch funcCallOrVar.Label {
	case "FunctionCall":
		HandleFunctionCall(&funcCallOrVar)
	case "VariableName":
		HandleVariableName(&funcCallOrVar)
	}
}

func HandleSequence(mt *MatchTree) {
	for i := 1; i < len(mt.Children); i++ {
		child := mt.Children[i]
		switch child.Label {
		case "VariableName":
			HandleVariableName(&child)
		case "FunctionCall":
			HandleFunctionCall(&child)
		}
	}
}

func HandleSequenceOfCharacters(mt *MatchTree) {
	//should never get this far
	return
}

func processVarBlockAssignmentStatement(mt *MatchTree) MatchTree {
	mtNameNode := mt.Children[0].Children[0]
	mtValueNode := mt.Children[2].Children[0]

	statement := MatchTree{IsValid: true, Value: mt.Value, Type: "Set", Label: "Statement"}
	name := MatchTree{IsValid: true, Value: mtNameNode.Value, Type: "Set", Label: "Name"}
	statement.Children = append(statement.Children, name)

	if mtValueNode.Label == "VariableName" {
		statement.Children = append(statement.Children, MatchTree{IsValid: true, Value: mtValueNode.Value, Type: "Set", Label: "Reference"})
	} else if mtValueNode.Label == "FunctionCall" {
		statement.Children = append(statement.Children, walkFunctionCall(&mtValueNode))
	}

	return statement
}

func walkFunctionCall(mt *MatchTree) MatchTree {
	mtNameNode := mt.Children[0]
	mtFunctionArgs := mt.Children[1]

	valueNode := MatchTree{IsValid: true, Value: mt.Value, Type: "Set", Label: mtNameNode.Value}

	for _, mtFuncArg := range mtFunctionArgs.Children {
		mtFuncArgChild := mtFuncArg.Children[0]
		mtFuncArgChildLabel := mtFuncArgChild.Label

		if strings.Contains(mtFuncArgChildLabel, ":") {
			mtFuncArgChildLabel = strings.Split(mtFuncArgChildLabel, ":")[1]
		}

		if mtFuncArgChildLabel == "string" {
			valueNode.Children = append(valueNode.Children, MatchTree{IsValid: true, Value: mtFuncArgChild.Value, Type: "Set", Label: "String"})
		} else if mtFuncArgChildLabel == "int" {
			valueNode.Children = append(valueNode.Children, MatchTree{IsValid: true, Value: mtFuncArgChild.Value, Type: "Set", Label: "Int"})
		} else if mtFuncArgChildLabel == "bool" {
			valueNode.Children = append(valueNode.Children, MatchTree{IsValid: true, Value: mtFuncArgChild.Value, Type: "Set", Label: "Bool"})
		} else if mtFuncArgChildLabel == "FunctionCall" {
			valueNode.Children = append(valueNode.Children, walkFunctionCall(&mtFuncArgChild))
		} else if mtFuncArgChildLabel == "VariableName" {
			valueNode.Children = append(valueNode.Children, MatchTree{IsValid: true, Value: mtFuncArgChild.Value, Type: "Set", Label: "Reference"})
		}
	}

	return valueNode
}

func parseFile(filepath string, statementz *MatchTree) {
	fileAsBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err.Error())
	}
	tree := Match(string(fileAsBytes), Golang)
	tree = tree.PruneToLabels()
	//fmt.Println(tree.ToMermaidDiagram())
	visit := func(mt *MatchTree) {
		switch mt.Label {
		case "VarBlockAssignmentStatement":
			statementz.Children = append(statementz.Children, processVarBlockAssignmentStatement(mt))
		}
	}
	tree.AcceptVisitor(visit)
}

func Test_Golang(t *testing.T) {
	t.Run("Golang", func(t *testing.T) {
		statementz := MatchTree{Type: "Set", Label: "Statements"}
		parseFile("./package_declaration.go", &statementz)

		//fmt.Println(statementz.ToMermaidDiagram())
	})
}
