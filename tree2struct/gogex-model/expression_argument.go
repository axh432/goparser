package gogex_model

import (
	"errors"
	"fmt"
	. "github.com/axh432/gogex"
)

type ExpressionArgument_t struct {
	ExpressionType           string
	ExpressionName           string
	SequenceCall             *SequenceCall_t
	SequenceOfCharactersCall *SequenceOfCharactersCall_t
	SetOfCharactersCall      *SetOfCharactersCall_t
	SetOfNotCharactersCall   *SetOfNotCharactersCall_t
	SetCall                  *SetCall_t
	RangeCall                *RangeCall_t
	LabelCall                *LabelCall_t
}

func NewExpressionArgument(mt *MatchTree) (*ExpressionArgument_t, error) {
	if len(mt.Labels) == 0 {
		return nil, errors.New("<ExpressionArgument> given match tree was not an ExpressionArgument: the provided matchtree has no labels")
	}
	if mt.Labels[0] != "name:ExpressionArgument" {
		return nil, fmt.Errorf("<ExpressionArgument> given match tree was not an ExpressionArgument: %s", mt.Labels[0])
	}
	if len(mt.Children) != 1 {
		return nil, errors.New("<ExpressionArgument> given match tree was not an ExpressionArgument: the match tree does not have exactly one child")
	}
	switch mt.Children[0].Labels[0] {
	case "name:ExpressionName":
		return &ExpressionArgument_t{ExpressionName: mt.Children[0].Value, ExpressionType: "ExpressionName"}, nil
	case "name:SequenceCall":
		seqCall, err := NewSequenceCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{SequenceCall: seqCall, ExpressionType: "SequenceCall"}, nil
	case "name:SequenceOfCharactersCall":
		seqOfCharCall, err := NewSequenceOfCharactersCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{SequenceOfCharactersCall: seqOfCharCall, ExpressionType: "SequenceOfCharactersCall"}, nil
	case "name:SetOfCharactersCall":
		setOfCharCall, err := NewSetOfCharactersCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{SetOfCharactersCall: setOfCharCall, ExpressionType: "SetOfCharactersCall"}, nil
	case "name:SetOfNotCharactersCall":
		setOfNotCharCall, err := NewSetOfNotCharactersCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{SetOfNotCharactersCall: setOfNotCharCall, ExpressionType: "SetOfNotCharactersCall"}, nil
	case "name:SetCall":
		setCall, err := NewSetCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{SetCall: setCall, ExpressionType: "SetCall"}, nil
	case "name:RangeCall":
		rangeCall, err := NewRangeCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{RangeCall: rangeCall, ExpressionType: "RangeCall"}, nil
	case "name:LabelCall":
		labelCall, err := NewLabelCall(&mt.Children[0])
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument\n%w", err)
		}
		return &ExpressionArgument_t{LabelCall: labelCall, ExpressionType: "LabelCall"}, nil
	}
	return nil, fmt.Errorf("<ExpressionArgument> unable to create ExpressionArgument: the child was not one of the possible options. The label was: %s", mt.Children[0].Labels[0])
}

func NewExpressionArguments(mt *MatchTree) ([]*ExpressionArgument_t, error) {
	if err := validateSliceTree(mt, "ExpressionArguments", "ExpressionArgument"); err != nil {
		return nil, err
	}
	var exps []*ExpressionArgument_t
	for _, child := range mt.Children {
		exp, err := NewExpressionArgument(&child)
		if err != nil {
			return nil, fmt.Errorf("<ExpressionArguments> unable to create ExpressionArguments\n%w", err)
		}
		exps = append(exps, exp)
	}
	return exps, nil
}
