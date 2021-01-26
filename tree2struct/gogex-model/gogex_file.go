package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
)

type GogexFile_t struct {
	PackageStatement *PackageStatement_t
	ImportStatement  *ImportStatement_t
	VarStatement     *VarStatement_t
}

func NewGogexFile(mt *MatchTree) (*GogexFile_t, error) {
	structName := "GogexFile"
	if err := validateStructTree(mt, structName, []string{"PackageStatement", "ImportStatement", "VarStatement"}); err != nil {
		return nil, err
	}
	packageStatement, err := NewPackageStatement(&mt.Children[0])
	if err != nil {
		return nil, fmt.Errorf("<GogexFile> unable to create %s. Issue with creating PackageStatement\n%w", structName, err)
	}
	importStatement, err := NewImportStatement(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<GogexFile> unable to create %s. Issue with creating ImportStatement\n%w", structName, err)
	}
	varStatement, err := NewVarStatement(&mt.Children[2])
	if err != nil {
		return nil, fmt.Errorf("<GogexFile> unable to create %s. Issue with creating VarStatement\n%w", structName, err)
	}
	return &GogexFile_t{packageStatement, importStatement, varStatement}, nil
}
