package goparser

import (
	_ "fmt"
	_ "strings"
)

/*Todo: A way to simplify this is to turn all whitespace into optionalWhitespaceBlocks
-This way it wont fail on bad go code but also wont fail on good go code and ultimately the point of this
is to parse go not validate it.*/
//Todo: We should organise the base level stuff? - they shouldn't have labels until they are used in a grammar.
//Todo: we should write functions for the logic that repeats.
//Todo: We should write filterable labels so that pruneToLabels() can offer different views.
//Todo: Move all the tests out of golang_grammar_test.go and into their own files
//Todo: start parsing the package declaration as that seems to be the simplest
