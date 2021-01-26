package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
)

type SetCall_t struct {
	SetKeyword          string
	ExpressionArguments []*ExpressionArgument_t
}

func NewSetCall(mt *MatchTree) (*SetCall_t, error) {
	structName := "SetCall"
	if err := validateStructTree(mt, structName, []string{"SetKeyword", "ExpressionArguments"}); err != nil {
		return nil, err
	}
	setKeyword := mt.Children[0].Value
	expressionArgument, err := NewExpressionArguments(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<SetCall> unable to create %s. Issue with creating ExpressionArguments\n%w", structName, err)
	}
	return &SetCall_t{setKeyword, expressionArgument}, nil
}
