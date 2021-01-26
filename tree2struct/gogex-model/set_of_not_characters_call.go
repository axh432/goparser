package gogex_model

import (
	. "github.com/axh432/gogex"
)

type SetOfNotCharactersCall_t struct {
	SetOfNotCharactersKeyword string
	Argument                  string
}

func NewSetOfNotCharactersCall(mt *MatchTree) (*SetOfNotCharactersCall_t, error) {
	if err := validateStructTree(mt, "SetOfNotCharactersCall", []string{"SetOfNotCharactersKeyword", "Argument"}); err != nil {
		return nil, err
	}
	return &SetOfNotCharactersCall_t{mt.Children[0].Value, mt.Children[1].Value}, nil
}
