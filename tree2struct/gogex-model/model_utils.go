package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
	"strings"
)

func validateStructTree(mt *MatchTree, expectedName string, expectedOffspring []string) error {
	if len(mt.Labels) == 0 {
		return fmt.Errorf("given match tree was not a %s: the provided matchtree has no labels")
	}
	if mt.Labels[0] != fmt.Sprintf("name:%s", expectedName) {
		return fmt.Errorf("<%s> given match tree was not a %s. The match tree's name was: %s", expectedName, expectedName, mt.Labels[0])
	}

	sb := strings.Builder{}
	for _, child := range mt.Children {
		sb.WriteString(" ")
		sb.WriteString(child.Labels[0])
	}
	childErr := fmt.Errorf("<%s> given match tree was not a %s: expected %d chldren for the given match tree: [%v] but instead found: [%s]",
		expectedName,
		expectedName,
		len(expectedOffspring),
		expectedOffspring,
		sb.String())

	if len(mt.Children) != len(expectedOffspring) {
		return childErr
	}
	for i, child := range mt.Children {
		if child.Labels[0] != "name:"+expectedOffspring[i] {
			return childErr
		}
	}

	return nil
}

func validateSliceTree(mt *MatchTree, expectedName string, expectedOffspring string) error {
	if len(mt.Labels) == 0 {
		return fmt.Errorf("<%s> given match tree was not a %s: the provided matchtree has no labels", expectedName, expectedName)
	}
	if mt.Labels[0] != fmt.Sprintf("name:%s", expectedName) {
		return fmt.Errorf("<%s> given match tree was not a %s: %s", expectedName, expectedName, mt.Labels[0])
	}
	for _, child := range mt.Children {
		if child.Labels[0] != "name:"+expectedOffspring {
			return fmt.Errorf("<%s> given match tree was not a %s: expected children to all be: [%s] but instead found: [%s]", expectedName, expectedName, expectedOffspring, child.Labels[0])
		}
	}
	return nil
}
