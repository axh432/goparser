package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
)

type ExpressionStatement_t struct {
	ExpressionName     string
	ExpressionArgument *ExpressionArgument_t
}

func NewExpressionStatement(mt *MatchTree) (*ExpressionStatement_t, error) {
	if err := validateStructTree(mt, "ExpressionStatement", []string{"ExpressionName", "ExpressionArgument"}); err != nil {
		return nil, err
	}
	expArg, err := NewExpressionArgument(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<ExpressionStatement> unable to create ExpressionStatement\n%w", err)
	}
	return &ExpressionStatement_t{mt.Children[0].Value, expArg}, nil
}

func NewExpressionStatements(mt *MatchTree) ([]*ExpressionStatement_t, error) {
	if err := validateSliceTree(mt, "Expressions", "ExpressionStatement"); err != nil {
		return nil, err
	}
	var exps []*ExpressionStatement_t
	for _, child := range mt.Children {
		exp, err := NewExpressionStatement(&child)
		if err != nil {
			return nil, fmt.Errorf("<ExpressionStatement> unable to create ExpressionStatement\n%w", err)
		}
		exps = append(exps, exp)
	}
	return exps, nil
}
