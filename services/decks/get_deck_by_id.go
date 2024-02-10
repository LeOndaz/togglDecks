package decks

import (
	"togglDecks/db"
	"togglDecks/models"
)

func GetDeckById(id string) (*models.Deck, error) {
	// two queries

	deck := &models.Deck{}

	if err := db.GetDB().Where("id = ?", id).Preload("Cards").First(deck).Error; err != nil {
		return nil, err
	}

	return deck, nil
}
