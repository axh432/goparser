package grammar

import (
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Values(t *testing.T) {
	t.Run("string value", func(t *testing.T) {
		require.True(t, Match(`"fmt \" line\""`, StringValue).IsValid)
		require.True(t, Match(`"github.com/axh432/gogex"`, StringValue).IsValid)
		require.True(t, Match("\"\"", StringValue).IsValid)
	})
}
