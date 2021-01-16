package modelvisitor

import (
	"fmt"
	"strings"
)

type StructDefinition struct {
	Name      string
	Type      string
	Variables map[string]string
}

type ArrayDefinition struct {
	Name string
	Type string
}

type PrimitiveDefinition struct {
	Name string
	Type string
}

type TypeDefinition interface {
	String() string
	GetType() string
}

func (sd StructDefinition) GetType() string {
	return sd.Type
}

func (ad ArrayDefinition) GetType() string {
	return ad.Type
}

func (pd PrimitiveDefinition) GetType() string {
	return pd.Type
}

func (sd StructDefinition) String() string {
	sb := strings.Builder{}
	sb.WriteString(sd.printDef())
	sb.WriteString("\n")
	sb.WriteString(sd.printConstructor())
	return sb.String()
}

func (sd StructDefinition) printDef() string {
	lw := LineWriter{}
	lw.WriteLine(0, fmt.Sprintf("type %s struct {", sd.Name))
	for Name, Type := range sd.Variables {
		lw.WriteLine(1, fmt.Sprintf("%s\t%s", Name, Type))
	}
	lw.WriteLine(0, "}")

	sb := strings.Builder{}

	sb.WriteString(lw.String())
	sb.WriteString("\n")
	sb.WriteString(sd.printConstructor())

	return lw.String()
}

type LineWriter struct {
	sb strings.Builder
}

func (lw *LineWriter) WriteLine(numberOfLeadingTabs int, line string) {
	lw.sb.WriteString(fmt.Sprintf("%s%s\n", createTabs(numberOfLeadingTabs), line))
}

func (lw *LineWriter) String() string {
	return lw.sb.String()
}

func createTabs(numberOfTabs int) string {
	sb := strings.Builder{}
	for i := 0; i < numberOfTabs; i++ {
		sb.WriteString("\t")
	}
	return sb.String()
}

func (sd StructDefinition) printConstructor() string {
	lw := LineWriter{}
	lw.WriteLine(0, fmt.Sprintf("func create%s(mt *MatchTree) (%s, error) {", sd.Name, sd.Name))
	lw.WriteLine(1, fmt.Sprintf("new%s := %s{}", sd.Name, sd.Name))
	lw.WriteLine(1, "for _, child := range mt.Children {")
	lw.WriteLine(2, "Name, Type := getNameAndType(child.Label)")
	lw.WriteLine(2, "switch Name {")

	for Name, Type := range sd.Variables {
		sd.printComplexVariable(Name, Type, &lw)
	}

	lw.WriteLine(2, "default:")
	lw.WriteLine(3, writeUnknownVariableError(sd.Name))
	lw.WriteLine(2, "}")
	lw.WriteLine(1, "}")
	lw.WriteLine(1, fmt.Sprintf("return %s, nil", sd.Name))
	lw.WriteLine(0, "}")

	return lw.String()
}

func writeUnknownToSliceError(arrayName string, arrayType string) string {
	errorLine := fmt.Sprintf(`return new%s, errors.New(fmt.Sprintf("The variable '[STRING_MARKER]' is unknown to the slice '%s'.", Name))`, arrayName, arrayType)
	return replaceStringMarkers(errorLine)
}

func writeUnknownVariableError(structName string) string {
	errorLine := fmt.Sprintf(`return new%s, errors.New(fmt.Sprintf("The variable '[STRING_MARKER]' is unknown to the struct '%s'.", Name))`, structName, structName)
	return replaceStringMarkers(errorLine)
}

func writeVariableTypeError(structName string, varName string, varType string) string {
	errorLine := fmt.Sprintf(`return new%s, errors.New(fmt.Sprintf("The type of the variable '%s' is expected to be: '%s'. But found: '[STRING_MARKER]' instead.", Type))`, structName, varName, varType)
	return replaceStringMarkers(errorLine)
}

func replaceStringMarkers(line string) string {
	return strings.ReplaceAll(line, "[STRING_MARKER]", "%s")
}

func (ad ArrayDefinition) printConstructor() string {
	lw := LineWriter{}
	elementType := strings.TrimPrefix(ad.Type, "[]")
	lw.WriteLine(0, fmt.Sprintf("func create%s(mt *MatchTree) (%s, error) {", ad.Name, ad.Type))
	lw.WriteLine(1, fmt.Sprintf("new%s := %s{}", ad.Name, ad.Type))
	lw.WriteLine(1, "for _, child := range mt.Children {")
	lw.WriteLine(2, "Name, Type := getNameAndType(child.Label)")
	lw.WriteLine(2, "switch Name {")
	lw.WriteLine(2, fmt.Sprintf("case \"%s\":", elementType))
	lw.WriteLine(3, fmt.Sprintf("if Type == \"%s\" {", elementType))
	lw.WriteLine(4, fmt.Sprintf("new%s, err := create%s(&child)", elementType, elementType))
	lw.WriteLine(4, "if err != nil {")
	lw.WriteLine(5, fmt.Sprintf("return new%s, err", ad.Name))
	lw.WriteLine(4, "}")
	lw.WriteLine(4, fmt.Sprintf("new%s = append(new%s, new%s)", ad.Name, ad.Name, elementType))
	lw.WriteLine(3, "}else{")
	lw.WriteLine(4, writeVariableTypeError(ad.Name, elementType, elementType))
	lw.WriteLine(3, "}")
	lw.WriteLine(2, "default:")
	lw.WriteLine(3, writeUnknownToSliceError(ad.Name, ad.Type))
	lw.WriteLine(2, "}")
	lw.WriteLine(1, "}")
	lw.WriteLine(1, fmt.Sprintf("return new%s, nil", ad.Name))
	lw.WriteLine(0, "}")
	return lw.String()
}

func (ad ArrayDefinition) String() string {
	return ad.printConstructor()
}

func (pd PrimitiveDefinition) printConstructor() string {
	lw := LineWriter{}
	lw.WriteLine(0, fmt.Sprintf("func create%s(mt *MatchTree) (string, error) {", pd.Name))
	lw.WriteLine(1, "return mt.Value, nil")
	lw.WriteLine(0, "}")
	return lw.String()
}

func (pd PrimitiveDefinition) String() string {
	return pd.printConstructor()
}

func (sd StructDefinition) printPrimitiveVariable(Name string, Type string, lw *LineWriter) {
	lw.WriteLine(2, fmt.Sprintf("case \"%s\":", Name))
	lw.WriteLine(3, fmt.Sprintf("if Type == \"%s\" {", Type))
	lw.WriteLine(4, fmt.Sprintf("new%s.%s = child.Value", sd.Name, Name))
	lw.WriteLine(3, "}else{")
	lw.WriteLine(4, writeVariableTypeError(sd.Name, Name, Type))
	lw.WriteLine(3, "}")
}

func (sd StructDefinition) printComplexVariable(Name string, Type string, lw *LineWriter) {
	lw.WriteLine(2, fmt.Sprintf("case \"%s\":", Name))
	lw.WriteLine(3, fmt.Sprintf("if Type == \"%s\" {", Type))
	lw.WriteLine(4, fmt.Sprintf("new%s, err := create%s(&child)", Name, Name))
	lw.WriteLine(4, "if err != nil {")
	lw.WriteLine(5, fmt.Sprintf("return new%s, err", sd.Name))
	lw.WriteLine(4, "}")
	lw.WriteLine(4, fmt.Sprintf("new%s.%s = new%s", sd.Name, Name, Name))
	lw.WriteLine(3, "}else{")
	lw.WriteLine(4, writeVariableTypeError(sd.Name, Name, Type))
	lw.WriteLine(3, "}")
}

func writeGetNameAndType() string {
	lw := LineWriter{}
	lw.WriteLine(0, "func getNameAndType(label string) (Name, Type string) {")
	lw.WriteLine(1, "if strings.Contains(label, \":\") {")
	lw.WriteLine(2, "nameAndType := strings.Split(label, \":\")")
	lw.WriteLine(2, "return nameAndType[0], nameAndType[1]")
	lw.WriteLine(1, "}")
	lw.WriteLine(1, "return label, label")
	lw.WriteLine(0, "}")
	return lw.String()
}
