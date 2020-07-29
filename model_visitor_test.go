package goparser

import (
	"fmt"
	"testing"
)

func TestStructDefinition_printDef(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sd := StructDefinition{
			Name: "PackageDeclaration",
			Variables: map[string]string{
				"PackageKeyword": "string",
				"PackageNames":   "[]string",
			},
		}

		fmt.Println(sd.printConstructor())
	})
}
