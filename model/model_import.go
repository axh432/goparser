package model

import (
	"errors"
	"fmt"
	. "github.com/axh432/gogex"
	"github.com/axh432/goparser"
)

type ImportStatement struct {
	ImportKeyword string
	Imports       []Import
}

type Import struct {
	ImportAccessType string
	ImportName       string
}

func createImportStatement(mt *MatchTree) (ImportStatement, error) {
	newImportStatement := ImportStatement{}
	for _, child := range mt.Children {
		Name, Type := goparser.getNameAndType(child.Label)
		switch Name {
		case "ImportKeyword":
			if Type == "string" {
				newImportStatement.ImportKeyword = child.Value
			} else {
				return newImportStatement, errors.New(fmt.Sprintf("The type of the variable 'ImportKeyword' is expected to be: '%s'. But found: '%s' instead.", "string", Type))
			}
		case "Imports":
			if Type == "[]Import" {
				imports, err := createImports(&child)
				if err != nil {
					return newImportStatement, err
				}
				newImportStatement.Imports = imports
			} else {
				return newImportStatement, errors.New(fmt.Sprintf("The type of the variable 'Imports' is expected to be: '%s'. But found: '%s' instead.", "[]string", Type))
			}
		default:
			return newImportStatement, errors.New(fmt.Sprintf("The variable '%s' is unknown to the struct '%s'.", Name, "ImportStatement"))
		}
	}
	return newImportStatement, nil
}

func createImports(mt *MatchTree) ([]Import, error) {
	newImports := []Import{}
	for _, child := range mt.Children {
		Name, Type := goparser.getNameAndType(child.Label)
		switch Name {
		case "Import":
			if Type == "Import" {
				newImport, err := createImport(&child)
				if err != nil {
					return newImports, err
				}
				newImports = append(newImports, newImport)
			} else {
				return newImports, errors.New(fmt.Sprintf("The type of the variable 'Import' is expected to be: '%s'. But found: '%s' instead.", "Import", Type))
			}
		default:
			return newImports, errors.New(fmt.Sprintf("The variable '%s' is unknown to the slice '%s'.", Name, "[]Imports"))
		}
	}
	return newImports, nil
}

func createImport(mt *MatchTree) (Import, error) {
	newImport := Import{}
	for _, child := range mt.Children {
		Name, Type := goparser.getNameAndType(child.Label)
		switch Name {
		case "ImportAccessType":
			if Type == "string" {
				newImport.ImportAccessType = child.Value
			} else {
				return newImport, errors.New(fmt.Sprintf("The type of the variable 'ImportAccessType' is expected to be: '%s'. But found: '%s' instead.", "string", Type))
			}
		case "ImportName":
			if Type == "string" {
				newImport.ImportName = child.Value
			} else {
				return newImport, errors.New(fmt.Sprintf("The type of the variable 'ImportName' is expected to be: '%s'. But found: '%s' instead.", "string", Type))
			}
		default:
			return newImport, errors.New(fmt.Sprintf("The variable '%s' is unknown to the struct 'Import'.", Name))
		}
	}
	return newImport, nil
}
