package tree2struct

import (
	. "github.com/axh432/goparser/tree2struct/gogex-model"
	"strings"
)

type TypeDefinition interface {
	Name() string
	Type() string
	PrintDef() string
	PrintConstructor() string
}

func writeUtilFunctions() string {
	sb := strings.Builder{}
	sb.WriteString(writeValidateStructTreeFunc())
	sb.WriteString("\n\n")
	sb.WriteString(writeValidateSliceTreeFunc())
	return sb.String()
}

var gogexBaseTypes = map[string]string{"Whitespace": "", "Number": "", "Letter": "", "Punctuation": "", "Symbol": ""}
var expressionArgs = map[string]*ExpressionArgument_t{}
var typeDefinitions = map[string]TypeDefinition{}

func writeDefinitions() string {
	sb := strings.Builder{}
	for _, def := range typeDefinitions {
		if isAStruct(def.Type()) {
			sb.WriteString(def.PrintDef())
			sb.WriteString("\n")
		}
		sb.WriteString(def.PrintConstructor())
		sb.WriteString("\n")
	}
	return sb.String()
}

func getTypeDefinitionsFromGogexFile(gogexFile *GogexFile_t) {
	addExpressionArgumentsToMap(gogexFile)
	traverseExpressionArguments(gogexFile)
}

func addExpressionArgumentsToMap(gogexFile *GogexFile_t) {
	for _, expression := range gogexFile.VarStatement.Expressions {
		expressionArgs[expression.ExpressionName] = expression.ExpressionArgument
	}
}

func getExpressionArgumentFromMap(expressionName string) *ExpressionArgument_t {
	if expArg, ok := expressionArgs[expressionName]; ok {
		return expArg
	}
	return nil
}

func traverseExpressionArguments(gogexFile *GogexFile_t) {
	for _, expression := range gogexFile.VarStatement.Expressions {
		traverseExpressionArgument(expression.ExpressionArgument)
	}
}

func traverseExpressionArgument(expArg *ExpressionArgument_t) []TypeDefinition {
	switch expArg.ExpressionType {
	case "ExpressionName":
		if newArg := getExpressionArgumentFromMap(expArg.ExpressionName); newArg != nil {
			return traverseExpressionArgument(newArg)
		}
		if _, ok := gogexBaseTypes[expArg.ExpressionName]; ok {
			break
		}
		panic("referenced name was not found in map: " + expArg.ExpressionName)
	case "SequenceCall":
		children := []TypeDefinition{}
		for _, newArg := range expArg.SequenceCall.ExpressionArguments {
			children = append(children, traverseExpressionArgument(newArg)...)
		}
		return children
	case "SetCall":
		children := []TypeDefinition{}
		for _, newArg := range expArg.SetCall.ExpressionArguments {
			children = append(children, traverseExpressionArgument(newArg)...)
		}
		return children
	case "RangeCall":
		return traverseExpressionArgument(expArg.RangeCall.ExpressionArgument)
	case "LabelCall":
		Type := removeQuotesAndTags(expArg.LabelCall.Labels[1])
		var typeDef TypeDefinition

		switch {
		case isAPrimitive(Type):
			typeDef = &PrimitiveDefinition{
				name_: removeQuotesAndTags(expArg.LabelCall.Labels[0]),
				type_: Type,
			}
		case isAnArray(Type):
			typeDef = &ArrayDefinition{
				name_:    removeQuotesAndTags(expArg.LabelCall.Labels[0]),
				type_:    Type,
				children: traverseExpressionArgument(expArg.LabelCall.ExpressionArgument),
			}
		case isAStruct(Type):
			typeDef = &StructDefinition{
				name_:    removeQuotesAndTags(expArg.LabelCall.Labels[0]),
				type_:    Type,
				children: traverseExpressionArgument(expArg.LabelCall.ExpressionArgument),
			}
		default:
			panic("unable to process label call: unknown type: " + Type)
		}

		typeDefinitions[typeDef.Name()] = typeDef
		return []TypeDefinition{typeDef}
	}
	return []TypeDefinition{}
}

func removeQuotesAndTags(str string) string {
	withoutName := strings.ReplaceAll(str, "name:", "")
	withoutType := strings.ReplaceAll(withoutName, "type:", "")
	return strings.ReplaceAll(withoutType, `"`, "")
}

func isAStruct(Type string) bool {
	return Type == "struct"
}

func isAnArray(Type string) bool {
	return strings.Contains(Type, "[]")
}

func isAPrimitive(Type string) bool {
	var goPrimitiveTypes = map[string]string{"float32": "", "int": "", "bool": "", "string": "", "rune": "", "byte": ""}
	_, ok := goPrimitiveTypes[Type]
	return ok
}
