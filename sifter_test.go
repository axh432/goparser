package goparser

import (
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

var notCurlyBrackets = Range(SetOfNotCharacters("{"), 1, -1)

func CurlyBrackets(iter *Iterator) MatchTree {
	return Sequence(SetOfCharacters("{"), Range(notCurlyBrackets, 1, -1), SetOfCharacters("}"))(iter)
}

var searchForCurlyBrackets = Range(Set(notCurlyBrackets, CurlyBrackets), 1, -1)

func Test_sifter(t *testing.T) {
	t.Run("2. Variables", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./go-by-example/methods.go")
		require.NoError(t, err)

		tree := Match(string(fileAsBytes), searchForCurlyBrackets)

		require.True(t, tree.IsValid)

		fmt.Println(tree.ToMermaidDiagram())
	})
}
