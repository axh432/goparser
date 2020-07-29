package goparser

/*

import (
	"fmt"
	"github.com/stretchr/testify/require"
	_ "strconv"
	"testing"
)

func TestGrammarComposer(t *testing.T) {

	t.Run("create expression for import", func(t *testing.T) {

		variants := []string{`import "fmt"`, `import ("fmt")`, "import \n(\"fmt\"\n\"strings\")", "import \n(\"fmt\"\n\"strings\"\n\"strconv\")"}

		expressionAsString := GrammarComposer(variants)

		fmt.Println(expressionAsString)
	})

	t.Run("create expression for states", func(t *testing.T) {

		variants := []string{`"fmt"`, `_"fmt"`, `."fmt"`, `.  "fmt"`, `_ 	"fmt"`}


		expectedExp := `Set(Sequence(String),Sequence(underscore,String),Sequence(dot,String),Sequence(dot,whitespaceNoNewLineBlock,String),Sequence(underscore,whitespaceNoNewLineBlock,String))`
		//expectedExp := `Sequence(Range(SetOfCharacters("_", "."), 0, 1), optionalWhitespaceNoNewLineBlock, String)`

		expressionAsString, tree := ComposeStatesNoOptionals(variants)

		require.Nil(t, tree)
		require.Equal(t, expectedExp, expressionAsString)
	})

	t.Run("create expression for lists", func(t *testing.T) {

		variants := []string{`"fmt"`, "\"fmt\"\n\"strings\""}

		expressionAsString, tree := ComposeList(variants)

		require.Nil(t, tree)

		fmt.Println(expressionAsString)

	})


}

func Test_grammarIntersect(t *testing.T) {

	t.Run("intersect grammar", func(t *testing.T) {
		grammar1 := "a,b,c,d,a,b,c,d,e,f,a,b,a,b,c,d"
		grammar2 := "c,d,e,f,c,d,e,f,c,d,e,f,c,d,e,f"

	})

}


func Test_grammarDeduplicate(t *testing.T) {

	t.Run("create expression for import", func(t *testing.T) {
		grammar := grammarDeduplicate("dot,whitespaceNoNewLineBlock,dot,String,String")
		fmt.Println(grammar)
		fmt.Println(grammarReduplicate(grammar))
	})

}

func Test_grammarIntersect(t *testing.T) {

	t.Run("create expression for import", func(t *testing.T) {
		intersection := grammarIntersect("underscore,String", "dot,String")
		fmt.Println(intersection)
	})

}*/
