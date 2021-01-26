package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
)

type VarStatement_t struct {
	VarKeyword  string
	Expressions []*ExpressionStatement_t
}

func NewVarStatement(mt *MatchTree) (*VarStatement_t, error) {
	if err := validateStructTree(mt, "VarStatement", []string{"VarKeyword", "Expressions"}); err != nil {
		return nil, err
	}
	exps, err := NewExpressionStatements(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<VarStatement> unable to create VarStatement\n%w", err)
	}
	return &VarStatement_t{mt.Children[0].Value, exps}, nil
}
