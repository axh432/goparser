package tree2struct

import (
	"fmt"
	"strings"
)

type LineWriter struct {
	sb strings.Builder
}

func (lw *LineWriter) WriteLine(numberOfLeadingTabs int, line string) {
	lw.sb.WriteString(fmt.Sprintf("%s%s\n", createTabs(numberOfLeadingTabs), line))
}

func (lw *LineWriter) String() string {
	return lw.sb.String()
}

func createTabs(numberOfTabs int) string {
	sb := strings.Builder{}
	for i := 0; i < numberOfTabs; i++ {
		sb.WriteString("\t")
	}
	return sb.String()
}
