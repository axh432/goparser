package tree2struct

import (
	"fmt"
	"strings"
)

type StructDefinition struct {
	name_    string
	type_    string
	children []TypeDefinition
}

func (sd StructDefinition) Name() string {
	return sd.name_
}

func (sd StructDefinition) Type() string {
	return sd.type_
}

func (sd StructDefinition) PrintDef() string {
	lw := LineWriter{}
	lw.WriteLine(0, fmt.Sprintf("type %s_t struct {", sd.Name()))
	for _, child := range sd.children {
		lw.WriteLine(1, fmt.Sprintf("%s\t%s", child.Name(), child.Type()))
	}
	lw.WriteLine(0, "}")
	return lw.String()
}

func (sd StructDefinition) PrintConstructor() string {
	lw := LineWriter{}
	childVarNames := []string{}
	lw.WriteLine(0, fmt.Sprintf("func New%s(mt *MatchTree) (*%s_t, error) {", sd.Name(), sd.Name()))
	lw.WriteLine(1, fmt.Sprintf(`if err := validateTree(mt, "%s"); err != nil {`, sd.Name()))
	lw.WriteLine(2, `return nil, err`)
	lw.WriteLine(1, `}`)
	for _, child := range sd.children {
		childVarName := lowerCaseFirstLetter(child.Name())
		childVarNames = append(childVarNames, childVarName)
		lw.WriteLine(1, fmt.Sprintf(`%s, err := New%s(&mt.Children[0])`, childVarName, child.Name()))
		lw.WriteLine(1, `if err != nil {`)
		lw.WriteLine(2, fmt.Sprintf(`return nil, fmt.Errorf("<%s> unable to be constructed. Issue with creating %s\n%s", err)`, sd.Name(), child.Name(), "%w"))
		lw.WriteLine(1, `}`)
	}
	lw.WriteLine(1, fmt.Sprintf(`return &%s_t%s, nil`, sd.Name(), formatChildVarNames(childVarNames)))
	lw.WriteLine(0, `}`)
	return lw.String()
}

func lowerCaseFirstLetter(str string) string {
	return strings.ToLower(string(str[0])) + str[1:]
}

func formatChildVarNames(varNames []string) string {
	asString := fmt.Sprintf("%v", varNames)
	withCommas := strings.ReplaceAll(asString, " ", ", ")
	leftBrace := strings.ReplaceAll(withCommas, "[", "{")
	rightBrace := strings.ReplaceAll(leftBrace, "]", "}")
	return rightBrace
}
