package grammar

import (
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Package_Declaration(t *testing.T) {

	t.Run("package declaration", func(t *testing.T) {
		tree := Match(`package somepackage`, PackageDeclaration)
		require.True(t, tree.IsValid)
	})
}
