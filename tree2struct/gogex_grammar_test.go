package tree2struct

import (
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func Test_GogexGrammar(t *testing.T) {
	t.Run("parse a gogex file", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./import_example.go")
		require.NoError(t, err)

		tree := Match(string(fileAsBytes), GogexFile)

		justLabels := tree.PruneToLabels()

		fmt.Println(justLabels.ToGraphVizDiagram())

		require.True(t, tree.IsValid)
	})

}

func TestRangeCall(t *testing.T) {
	t.Run("given a range of named expressionArgs return a valid parse tree", func(t *testing.T) {
		tree := Match("Range(BackTick, 1, 2)", RangeCall)
		labels := tree.PruneToLabels()
		fmt.Println(labels.ToGraphVizDiagram())
		require.True(t, tree.IsValid)
	})
}

func TestSequenceCall(t *testing.T) {
	t.Run("given a sequence of named expressionArgs return a valid parse tree", func(t *testing.T) {
		tree := Match("Sequence(BackTick, Something, BackTick)", SequenceCall)
		fmt.Println(tree.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})

	t.Run("given a sequence of nested sequences return a valid parse tree", func(t *testing.T) {
		tree := Match("Sequence(BackTick, Sequence(Sequence(Something), Something, Darkside), BackTick)", SequenceCall)
		fmt.Println(tree.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})

	t.Run("given a sequence of nested SequenceOfCharacters return a valid parse tree", func(t *testing.T) {
		tree := Match("Sequence(BackTick, SequenceOfCharacters(\"Something\"), BackTick)", SequenceCall)
		fmt.Println(tree.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})

	t.Run("given a sequence of nested SetOfCharacters return a valid parse tree", func(t *testing.T) {
		tree := Match("Sequence(BackTick, SetOfCharacters(\"'\"), BackTick)", SequenceCall)
		fmt.Println(tree.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})

	t.Run("given a sequence of nested SetOfNotCharacters return a valid parse tree", func(t *testing.T) {
		tree := Match("Sequence(BackTick, SetOfNotCharacters(\"'\"), BackTick)", SequenceCall)
		fmt.Println(tree.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})

	t.Run("given a SetOfNotCharacters return a valid parse tree", func(t *testing.T) {
		tree := Match("SetOfNotCharacters(\"'\")", SetOfNotCharactersCall)
		fmt.Println(tree.ToMermaidDiagram())
		require.True(t, tree.IsValid)
	})
}

func visualizeTree(tree MatchTree) {

	baseDir := "/Personal-Projects/mermaid-test"

	err := ioutil.WriteFile(baseDir+"/input.gv", []byte(tree.ToGraphVizDiagram()), 0)

	if err != nil {
		panic(err)
	}

}
