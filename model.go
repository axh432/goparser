package goparser

import (
	"fmt"
	"strings"
)

type Model interface {
	toString() string
}

type mString struct {
	String string
}

func (s *mString) toString() string {
	return s.String
}

type mFunctionCall struct {
	PackageName string
	Name        string
	Parameters  []Model
}

func (fc *mFunctionCall) toString() string {
	sb := strings.Builder{}

	if fc.PackageName != "" {
		sb.WriteString(fc.PackageName)
		sb.WriteString(".")
	}

	sb.WriteString(fc.Name)
	sb.WriteString("(")

	lenParam := len(fc.Parameters)

	if lenParam == 1 {
		sb.WriteString(fc.Parameters[0].toString())
	} else if lenParam > 1 {
		sb.WriteString(fc.Parameters[0].toString())
		for i := 1; i < lenParam; i++ {
			sb.WriteString(", ")
			sb.WriteString(fc.Parameters[i].toString())
		}
	}

	sb.WriteString(")")
	return sb.String()
}

type mFunction struct {
	Name       string
	Statements []Model
}

func (f *mFunction) toString() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("func %s() {", f.Name))
	for _, st := range f.Statements {
		sb.WriteString("\n\t")
		sb.WriteString(st.toString())
	}
	sb.WriteString("\n")
	sb.WriteString("}")
	return sb.String()
}

type mImport struct {
	importAccessType string
	importName       string
}

type mImportDeclaration struct {
	imports []mImport
}

func (i *mImportDeclaration) toString() string {

	lenImports := len(i.imports)

	if lenImports < 1 {
		return ""
	} else if lenImports == 1 {
		return fmt.Sprintf("import %s %s", i.imports[0].importAccessType, i.imports[0].importName)
	}

	sb := strings.Builder{}
	sb.WriteString("(")

	for _, name := range i.imports {
		sb.WriteString("\n\t")
		sb.WriteString(name)
	}

	sb.WriteString("\n")
	sb.WriteString(")")

	return sb.String()
}

type mPackage struct {
	Name string
}

func (p *mPackage) toString() string {
	return fmt.Sprintf("package %s", p.Name)
}
