package tree2struct

import "fmt"

type PrimitiveDefinition struct {
	name_ string
	type_ string
}

func (pd PrimitiveDefinition) Name() string {
	return pd.name_
}

func (pd PrimitiveDefinition) Type() string {
	return pd.type_
}

func (pd PrimitiveDefinition) PrintDef() string {
	return ""
}

func (pd PrimitiveDefinition) PrintConstructor() string {
	lw := LineWriter{}
	lw.WriteLine(0, fmt.Sprintf("func New%s(mt *MatchTree) (%s, error) {", pd.Name(), pd.Type()))
	lw.WriteLine(1, fmt.Sprintf(`if err := validateTree(mt, "%s"); err != nil {`, pd.Name()))
	lw.WriteLine(2, `return nil, err`)
	lw.WriteLine(1, `}`)
	typeErrorStr := fmt.Sprintf(`return nil, fmt.Errorf("<%s> unable to be constructed. Unable to parse value %s\n%s", err)`, pd.Name(), pd.Type(), "%w")
	switch pd.Type() {
	case "string":
		lw.WriteLine(1, `return mt.Value, nil`)
	case "int":
		lw.WriteLine(1, `value, err := strconv.ParseInt(mt.Value, 10, 32)`)
		lw.WriteLine(1, `if err != nil {`)
		lw.WriteLine(2, typeErrorStr)
		lw.WriteLine(1, `}`)
		lw.WriteLine(1, `return int(value), nil`)
	case "float32":
		lw.WriteLine(1, `value, err := strconv.ParseFloat(mt.Value, 32)`)
		lw.WriteLine(1, `if err != nil {`)
		lw.WriteLine(2, typeErrorStr)
		lw.WriteLine(1, `}`)
		lw.WriteLine(1, `return float32(value), nil`)
	case "bool":
		lw.WriteLine(1, `value, err := strconv.ParseBool(mt.Value)`)
		lw.WriteLine(1, `if err != nil {`)
		lw.WriteLine(2, typeErrorStr)
		lw.WriteLine(1, `}`)
		lw.WriteLine(1, `return value, nil`)
	case "rune":
		lw.WriteLine(1, `value, err := strconv.ParseInt(mt.Value, 10, 32)`)
		lw.WriteLine(1, `if err != nil {`)
		lw.WriteLine(2, typeErrorStr)
		lw.WriteLine(1, `}`)
		lw.WriteLine(1, `return rune(value), nil`)
	case "byte":
		lw.WriteLine(1, `value, err := strconv.ParseInt(mt.Value, 10, 16)`)
		lw.WriteLine(1, `if err != nil {`)
		lw.WriteLine(2, typeErrorStr)
		lw.WriteLine(1, `}`)
		lw.WriteLine(1, `return byte(value), nil`)
	}
	lw.WriteLine(0, `}`)
	return lw.String()
}
