package gogex_model

import (
	. "github.com/axh432/gogex"
)

type PackageStatement_t struct {
	PackageKeyword string
	PackageName    string
}

func NewPackageStatement(mt *MatchTree) (*PackageStatement_t, error) {
	if err := validateStructTree(mt, "PackageStatement", []string{"PackageKeyword", "PackageName"}); err != nil {
		return nil, err
	}
	return &PackageStatement_t{mt.Children[0].Value, mt.Children[1].Value}, nil
}
