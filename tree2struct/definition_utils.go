package tree2struct

func writeValidateSliceTreeFunc() string {
	lw := LineWriter{}
	lw.WriteLine(0, `func validateSliceTree(mt *MatchTree, expectedName string, expectedOffspring string) error {`)
	lw.WriteLine(1, `if len(mt.Labels) == 0 {`)
	lw.WriteLine(2, `return fmt.Errorf("given match tree was not a %s: the provided matchtree has no labels")`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, `if mt.Labels[0] != fmt.Sprintf("name:%s", expectedName) {`)
	lw.WriteLine(2, `return fmt.Errorf("<%s> given match tree was not a %s. The match tree's name was: %s", expectedName, expectedName, mt.Labels[0])`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, `for _, child := range mt.Children {`)
	lw.WriteLine(2, `if child.Labels[0] != "name:"+expectedOffspring {`)
	lw.WriteLine(3, `return fmt.Errorf("<%s> given match tree was not a %s: expected children to all be: [%s] but instead found: [%s]", expectedName, expectedName, expectedOffspring, child.Labels[0])`)
	lw.WriteLine(2, `}`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, `return nil`)
	lw.WriteLine(0, `}`)
	return lw.String()
}

func writeValidateStructTreeFunc() string {
	lw := LineWriter{}
	lw.WriteLine(0, `func validateTree(mt *MatchTree, expectedName string) error {`)
	lw.WriteLine(1, `if len(mt.Labels) == 0 {`)
	lw.WriteLine(2, `return fmt.Errorf("given match tree was not a %s: the provided matchtree has no labels")`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, `if mt.Labels[0] != fmt.Sprintf("name:%s", expectedName) {`)
	lw.WriteLine(2, `return fmt.Errorf("<%s> given match tree was not a %s. The match tree's name was: %s", expectedName, expectedName, mt.Labels[0])`)
	lw.WriteLine(1, `}`)
	lw.WriteLine(1, `return nil`)
	lw.WriteLine(0, `}`)
	return lw.String()
}
