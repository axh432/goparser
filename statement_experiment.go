package goparser

import "github.com/axh432/goparser/grammar/modelvisitor"

type CodeStatement func(numberOfLeadingTabs int) string

func create_codeBlockStatement(header string, footer string, statements ...CodeStatement) CodeStatement {
	return func(numberOfLeadingTabs int) string {
		lw := modelvisitor.LineWriter{}
		lw.WriteLine(numberOfLeadingTabs, header)
		for _, statement := range statements {
			lw.WriteLine(0, statement(numberOfLeadingTabs+1))
		}
		if footer != "" {
			lw.WriteLine(numberOfLeadingTabs, footer)
		}
		return lw.String()
	}
}

func create_singleLineStatement(functionCall string) CodeStatement {
	return func(numberOfLeadingTabs int) string {
		lw := modelvisitor.LineWriter{}
		lw.WriteLine(numberOfLeadingTabs, functionCall)
		return lw.String()
	}
}
