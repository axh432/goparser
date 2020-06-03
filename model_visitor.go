package goparser

import (
	"fmt"
	. "github.com/axh432/gogex"
	"strconv"
	"strings"
)

var ModelVisitor = func(mt *MatchTree) {
	sb := strings.Builder{}
	if mt.Label != "" {
		sb.WriteString(mt.Label)
		sb.WriteString("-")
	}
	sb.WriteString(strconv.Quote(mt.Value))
	fmt.Println(sb.String())
}

var ImportVisitor = func(mt *MatchTree) {
	sb := strings.Builder{}
	if mt.Label != "" {
		sb.WriteString(mt.Label)
		sb.WriteString("-")
	}
	sb.WriteString(strconv.Quote(mt.Value))
	fmt.Println(sb.String())
}
