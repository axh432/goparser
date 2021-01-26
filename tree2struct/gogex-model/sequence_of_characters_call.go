package gogex_model

import (
	. "github.com/axh432/gogex"
)

type SequenceOfCharactersCall_t struct {
	SequenceOfCharactersKeyword string
	Argument                    string
}

func NewSequenceOfCharactersCall(mt *MatchTree) (*SequenceOfCharactersCall_t, error) {
	if err := validateStructTree(mt, "SequenceOfCharactersCall", []string{"SequenceOfCharactersKeyword", "Argument"}); err != nil {
		return nil, err
	}
	return &SequenceOfCharactersCall_t{mt.Children[0].Value, mt.Children[1].Value}, nil
}
