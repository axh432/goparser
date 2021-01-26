package tree2struct

import "fmt"

type ArrayDefinition struct {
	name_    string
	type_    string
	children []TypeDefinition
}

func (ad ArrayDefinition) Name() string {
	return ad.name_
}

func (ad ArrayDefinition) Type() string {
	return ad.type_
}

func (ad ArrayDefinition) PrintDef() string {
	return ""
}

func (ad ArrayDefinition) PrintConstructor() string {
	lw := LineWriter{}
	child := ad.children[0]
	lw.WriteLine(0, fmt.Sprintf(`New%s(mt *MatchTree) ([]*%s_t, error) {`, ad.Name(), child.Name()))
	lw.WriteLine(1, fmt.Sprintf(`if err := validateSliceTree(mt, "%s", "%s"); err != nil {`, ad.Name(), child.Name()))
	lw.WriteLine(2, `return nil, err`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, fmt.Sprintf(`var exps []%s_t`, child.Name()))
	lw.WriteLine(1, `for _, child := range mt.Children {`)
	lw.WriteLine(2, fmt.Sprintf(`exp, err := New%s(&child)`, child.Name()))
	lw.WriteLine(2, `if err != nil {`)
	lw.WriteLine(3, fmt.Sprintf(`return nil, fmt.Errorf("<%s> unable to be constructed\n%s", err)`, ad.Name(), "%w"))
	lw.WriteLine(2, `}`)
	lw.WriteLine(2, `exps = append(exps, exp)`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, `return exps, nil`)
	lw.WriteLine(0, `}`)
	return lw.String()
}
