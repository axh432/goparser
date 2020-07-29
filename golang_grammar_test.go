package goparser

import (
	"fmt"
	. "github.com/axh432/gogex"
	. "github.com/axh432/goparser/grammar"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func Test_go_grammar(t *testing.T) {

	t.Run("var assign statement", func(t *testing.T) {
		require.True(t, Match(`var a = "initial"`, VarAssignmentStatement).IsValid)
		require.True(t, Match(`var d = true"`, VarAssignmentStatement).IsValid)
		require.True(t, Match(`var b int = 1`, VarAssignmentStatement).IsValid)
		require.True(t, Match(`var b, c int = 1, 2"`, VarAssignmentStatement).IsValid)
		require.True(t, Match(`var b = strconv.toInt("12")`, VarAssignmentStatement).IsValid)
	})

	t.Run("var declaration statement", func(t *testing.T) {
		require.True(t, Match(`var e int`, VarStatement).IsValid)
	})

	t.Run("var block statement", func(t *testing.T) {
		tree := Match("var (\ntypeName     = variableName)", VarBlock)

		fmt.Println(tree.ToMermaidDiagram())

		require.True(t, tree.IsValid)

		require.True(t, Match("var (\ntypeName     = Label(name, \"TypeName:string\")\n)", VarBlock).IsValid)
	})

	t.Run("assign statement", func(t *testing.T) {
		require.True(t, Match(`f := "apple"`, AssignmentStatement).IsValid)
	})

	t.Run("function call", func(t *testing.T) {
		require.True(t, Match(`Println(a)`, FunctionCall).IsValid)
		require.True(t, Match(`fmt.Println(a)`, FunctionCall).IsValid)
		require.True(t, Match(`fmt.Println(b, c)`, FunctionCall).IsValid)
		require.True(t, Match(`fmt.Println(fmt.Sprintf("hello %d", d), c)`, FunctionCall).IsValid)
	})

	t.Run("function body", func(t *testing.T) {
		require.True(t, Match("{ \tvar b, c int = 1, 2\n\tfmt.Println(b, c) }", FunctionBody).IsValid)
		require.True(t, Match("{ \tfmt.Println(b, c) }", FunctionBody).IsValid)

		require.True(t, Match("{\n}", FunctionBody).IsValid)
		require.True(t, Match("{}", FunctionBody).IsValid)
		require.True(t, Match("{\t}", FunctionBody).IsValid)

		require.True(t, Match("{\n\n\t}", FunctionBody).IsValid)
		require.True(t, Match("{\n\n\tvar a = \"initial\"\n\t}", FunctionBody).IsValid)

		require.True(t, Match("{\n\n\tvar a = \"initial\"\n\tfmt.Println(a)\n\n\tvar b, c int = 1, 2\n\tfmt.Println(b, c)\n\n\tvar d = true\n\tfmt.Println(d)}", FunctionBody).IsValid)
	})

	t.Run("function", func(t *testing.T) {
		tree := Match("func special() {\n\tvar a = \"initial\"\n}", FunctionDeclaration)
		require.True(t, tree.IsValid)
	})

	t.Run("1. Hello World", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./go-by-example/hello_world.go")
		require.NoError(t, err)

		tree := Match(string(fileAsBytes), Golang)

		fmt.Println(tree.PruneToLabels().ToMermaidDiagram())

		require.True(t, tree.IsValid)
	})

	t.Run("2. Variables", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./go-by-example/variables.go")
		require.NoError(t, err)

		tree := Match(string(fileAsBytes), Golang)

		require.True(t, tree.IsValid)
	})

	t.Run("3. Can it parse itself???", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./grammar/package_declaration.go")
		require.NoError(t, err)

		tree := Match(string(fileAsBytes), Golang)

		fmt.Println(tree.ToGraphVizDiagram())

		require.True(t, tree.IsValid)
	})
}
