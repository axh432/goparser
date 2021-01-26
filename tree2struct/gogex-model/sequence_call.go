package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
)

type SequenceCall_t struct {
	SequenceKeyword     string
	ExpressionArguments []*ExpressionArgument_t
}

func NewSequenceCall(mt *MatchTree) (*SequenceCall_t, error) {
	structName := "SequenceCall"
	if err := validateStructTree(mt, structName, []string{"SequenceKeyword", "ExpressionArguments"}); err != nil {
		return nil, err
	}
	sequenceKeyword := mt.Children[0].Value
	expressionArgument, err := NewExpressionArguments(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<SequenceCall> unable to create %s. Issue with creating ExpressionArguments\n%w", structName, err)
	}
	return &SequenceCall_t{sequenceKeyword, expressionArgument}, nil
}
