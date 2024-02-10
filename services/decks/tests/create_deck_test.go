package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"togglDecks/models"
	"togglDecks/services/decks"
)

func TestInitEmptyDeck(t *testing.T) {
	emptyDeck := &models.Deck{
		Name:     "Some empty deck for testing",
		Shuffled: false,
	}

	err := decks.InitEmptyDeck(emptyDeck)

	assert.Nil(t, err)
	assert.Equal(t, len(emptyDeck.Cards), 0, "Deck.Cards is not empty")
	assert.Equal(t, emptyDeck.Shuffled, false, "Deck.Shuffled can't be true")
	assert.NotNil(t, emptyDeck.ID, "Deck has no ID, InitEmptyDeck should have save to DB")
}
