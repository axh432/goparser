package gogex_model

import (
	"fmt"
	. "github.com/axh432/gogex"
)

type LabelCall_t struct {
	LabelKeyword       string
	ExpressionArgument *ExpressionArgument_t
	Labels             []string
}

func NewLabelCall(mt *MatchTree) (*LabelCall_t, error) {
	structName := "LabelCall"
	if err := validateStructTree(mt, structName, []string{"LabelKeyword", "ExpressionArgument", "Labels"}); err != nil {
		return nil, err
	}
	labelKeyword := mt.Children[0].Value
	expressionArgument, err := NewExpressionArgument(&mt.Children[1])
	if err != nil {
		return nil, fmt.Errorf("<LabelCall> unable to create %s. Issue creating ExpressionArgument\n%w", structName, err)
	}
	labels, err := NewLabels(&mt.Children[2])
	if err != nil {
		return nil, fmt.Errorf("<LabelCall> unable to create %s. Issue creating Labels\n%w", structName, err)
	}
	return &LabelCall_t{labelKeyword, expressionArgument, labels}, nil
}

func NewLabels(mt *MatchTree) ([]string, error) {
	if err := validateSliceTree(mt, "Labels", "Label"); err != nil {
		return nil, err
	}
	var labels []string
	for _, child := range mt.Children {
		labels = append(labels, child.Value)
	}
	return labels, nil
}
