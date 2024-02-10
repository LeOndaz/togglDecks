package decks

import (
	"togglDecks/db"
	"togglDecks/models"
)

func InitEmptyDeck(deck *models.Deck) error {
	err := ValidateDeck(deck)

	if err != nil {
		return err
	}

	if err := db.GetDB().Create(deck).Error; err != nil {
		return err
	}

	return nil
}
