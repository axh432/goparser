package gogexwalker

import (
	"fmt"
	. "github.com/axh432/gogex"
)

func Walk(mt *MatchTree) {
	fmt.Printf("%s\n", mt.Labels)
	for _, child := range mt.Children {
		Walk(&child)
	}
}
