package gogex_model

import (
	. "github.com/axh432/gogex"
)

type ImportStatement_t struct {
	ImportKeyword        string
	ImportAccessOperator string
	ImportURL            string
}

func NewImportStatement(mt *MatchTree) (*ImportStatement_t, error) {
	if err := validateStructTree(mt, "ImportStatement", []string{"ImportKeyword", "ImportAccessOperator", "ImportURL"}); err != nil {
		return nil, err
	}
	return &ImportStatement_t{mt.Children[0].Value, mt.Children[1].Value, mt.Children[2].Value}, nil
}
