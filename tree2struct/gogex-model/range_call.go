package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
	"strconv"
)

type RangeCall_t struct {
	RangeKeyword       string
	ExpressionArgument *ExpressionArgument_t
	MinValue           int
	MaxValue           int
}

func NewRangeCall(mt *MatchTree) (*RangeCall_t, error) {
	structName := "RangeCall"
	if err := validateStructTree(mt, structName, []string{"RangeKeyword", "ExpressionArgument", "MinValue", "MaxValue"}); err != nil {
		return nil, err
	}
	rangeKeyword := mt.Children[0].Value
	expressionArgument, err := NewExpressionArgument(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<RangeCall> unable to create %s. Issue with creating ExpressionArgument\n%w", structName, err)
	}
	minValue, err := strconv.Atoi(mt.Children[2].Value)
	if err != nil {
		return nil, fmt.Errorf("<RangeCall> unable to create %s. Issue with creating MinValue. Attempted to convert string '%s'\n%w", structName, mt.Children[2].Value, err)
	}
	maxValue, err := strconv.Atoi(mt.Children[3].Value)
	if err != nil {
		return nil, fmt.Errorf("<RangeCall> unable to create %s. Issue with creating MaxValue. Attempted to convert string '%s'\n%w", structName, mt.Children[3].Value, err)
	}
	return &RangeCall_t{rangeKeyword, expressionArgument, minValue, maxValue}, nil
}
