package common

import (
	"errors"
	"fmt"
)

type Error struct {
	Code     string
	Message  string
	Metadata any
}

const (
	ValidationError = "validation_error"
	UnknownError    = "unknown_error"

	InvalidQueryParameter = "invalid_query_parameter"
	InvalidPathParameter  = "invalid_path_parameter"

	CardNotFound = "card_not_found"

	DeckNotFound     = "deck_not_found"
	DeckFailedToDraw = "deck_failed_to_draw"
)

// TODO, reflect get method name?

func MustBeCalledOnNonNil(method string) error {
	return errors.New(
		fmt.Sprintf("%s must be called on an initialized instance, have you called it on a nil pointer", method),
	)
}
