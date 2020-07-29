package grammar

import (
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Import(t *testing.T) {
	t.Run("import", func(t *testing.T) {
		require.True(t, Match(`import "fmt"`, ImportDeclaration).IsValid)
		require.True(t, Match(`import ("fmt")`, ImportDeclaration).IsValid)
		require.True(t, Match("import (\"fmt\"\n\"strings\")", ImportDeclaration).IsValid)

		tree := Match("import (.\"fmt\"\n_\"strings\")", ImportDeclaration)

		tree = tree.PruneToLabels()

		require.True(t, tree.IsValid)
	})
}
