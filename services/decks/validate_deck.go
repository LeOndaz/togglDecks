package decks

import (
	"errors"
	"togglDecks/models"
)

func ValidateDeck(deck *models.Deck) error {
	if deck.Name == "" {
		return errors.New("deck must have a name")
	}

	return nil
}
