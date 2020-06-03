package goparser

import (
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func Test_go_grammar(t *testing.T) {

	t.Run("test function signature", func(t *testing.T) {
		require.True(t, Match("func copy()", functionSignature).IsValid)
		require.True(t, Match("func copy	()", functionSignature).IsValid)
		require.True(t, Match("func copy(	)", functionSignature).IsValid)
		require.True(t, Match("func copy()()", functionSignature).IsValid)
		require.True(t, Match("func copy(int left)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, \nint right)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right, float up)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) int", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (int, int)", functionSignature).IsValid)
		require.True(t, Match("func copy(int left, int right) (dave int, sedric int)", functionSignature).IsValid)
	})

	t.Run("strings", func(t *testing.T) {
		require.True(t, Match(`"fmt \" line\""`, String).IsValid)
		require.True(t, Match("\"\"", String).IsValid)
	})

	t.Run("import", func(t *testing.T) {
		require.True(t, Match(`import "fmt"`, importStatement).IsValid)
		require.True(t, Match(`import ("fmt")`, importStatement).IsValid)
		require.True(t, Match("import (\"fmt\"\n\"strings\")", importStatement).IsValid)

		tree := Match("import (.\"fmt\"\n_\"strings\")", importStatement)

		tree = tree.PruneToLabels()

		tree.AcceptVisitor(ModelVisitor)

		//fmt.Println(tree.PruneToLabels().ToMermaidDiagram())
	})

	t.Run("package", func(t *testing.T) {
		require.True(t, Match(`package somepackage`, packageDeclaration).IsValid)
	})

	t.Run("var assign statement", func(t *testing.T) {
		require.True(t, Match(`var a = "initial"`, varAssignStatement).IsValid)
		require.True(t, Match(`var d = true"`, varAssignStatement).IsValid)
		require.True(t, Match(`var b int = 1`, varAssignStatement).IsValid)
		require.True(t, Match(`var b, c int = 1, 2"`, varAssignStatement).IsValid)
		require.True(t, Match(`var b = strconv.toInt("12")`, varAssignStatement).IsValid)
	})

	t.Run("var declaration statement", func(t *testing.T) {
		require.True(t, Match(`var e int`, varStatement).IsValid)
	})

	t.Run("assign statement", func(t *testing.T) {
		require.True(t, Match(`f := "apple"`, assignStatement).IsValid)
	})

	t.Run("function call", func(t *testing.T) {
		require.True(t, Match(`Println(a)`, functionCall).IsValid)
		require.True(t, Match(`fmt.Println(a)`, functionCall).IsValid)
		require.True(t, Match(`fmt.Println(b, c)`, functionCall).IsValid)
		require.True(t, Match(`fmt.Println(fmt.Sprintf("hello %d", d), c)`, functionCall).IsValid)
	})

	t.Run("function body", func(t *testing.T) {
		require.True(t, Match("{ \tvar b, c int = 1, 2\n\tfmt.Println(b, c) }", functionBody).IsValid)
		require.True(t, Match("{ \tfmt.Println(b, c) }", functionBody).IsValid)

		require.True(t, Match("{\n}", functionBody).IsValid)
		require.True(t, Match("{}", functionBody).IsValid)
		require.True(t, Match("{\t}", functionBody).IsValid)

		require.True(t, Match("{\n\n\t}", functionBody).IsValid)
		require.True(t, Match("{\n\n\tvar a = \"initial\"\n\t}", functionBody).IsValid)

		require.True(t, Match("{\n\n\tvar a = \"initial\"\n\tfmt.Println(a)\n\n\tvar b, c int = 1, 2\n\tfmt.Println(b, c)\n\n\tvar d = true\n\tfmt.Println(d)}", functionBody).IsValid)
	})

	t.Run("function", func(t *testing.T) {
		tree := Match("func special() {\n\tvar a = \"initial\"\n}", functionDeclaration)
		require.True(t, tree.IsValid)
	})

	t.Run("1. Hello World", func(t *testing.T) {
		fileAsBytes, err := ioutil.ReadFile("./go-by-example/model.go")
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
}
