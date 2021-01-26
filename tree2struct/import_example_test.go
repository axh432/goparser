package tree2struct

import (
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ImportExample(t *testing.T) {
	t.Run("parse a gogex file and convert to struct", func(t *testing.T) {
		tree := Match(`import . "github.com/axh432/gogex"`, ImportStatement)
		require.True(t, tree.IsValid)
		justLabels := tree.PruneToLabels()
		fmt.Println(justLabels.ToMermaidDiagram())
	})

}
