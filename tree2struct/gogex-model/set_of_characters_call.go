package gogex_model

import . "github.com/axh432/gogex"

type SetOfCharactersCall_t struct {
	SetOfCharactersKeyword string
	Argument               string
}

func NewSetOfCharactersCall(mt *MatchTree) (*SetOfCharactersCall_t, error) {
	if err := validateStructTree(mt, "SetOfCharactersCall", []string{"SetOfCharactersKeyword", "Argument"}); err != nil {
		return nil, err
	}
	return &SetOfCharactersCall_t{mt.Children[0].Value, mt.Children[1].Value}, nil
}
