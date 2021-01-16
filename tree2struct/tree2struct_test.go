package tree2struct

import (
	"fmt"
	. "github.com/axh432/gogex"
	"testing"
)

func Test_Tree2Struct(t *testing.T) {
	t.Run("parse an import command to a go struct", func(t *testing.T) {
		tree := Match(`import . "github.com/axh432/gogex"`, ImportStatement)

		justLabels := tree.PruneToLabels()

		fmt.Println(justLabels.ToMermaidDiagram())
	})

}
