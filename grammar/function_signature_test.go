package grammar

import (
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Function_Signature(t *testing.T) {
	t.Run("test function signature", func(t *testing.T) {
		require.True(t, Match("func copy()", FunctionSignature).IsValid)
		require.True(t, Match("func copy	()", FunctionSignature).IsValid)
		require.True(t, Match("func copy(	)", FunctionSignature).IsValid)
		require.True(t, Match("func copy()()", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, \nint right)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right, float up)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) int", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int, int)", FunctionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int, sedric int)", FunctionSignature).IsValid)
	})
}
