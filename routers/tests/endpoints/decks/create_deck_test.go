package decks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"togglDecks/models"
	"togglDecks/testutils"
)

func TestCreateValidDeck(t *testing.T) {
	const deckName = "TestDeck"

	recorder := testutils.MockRequest(t, Router, testutils.MockApiCall{
		Method: "POST",
		Path:   "/decks",
		Body: models.Deck{
			Name: deckName,
		},
		QueryParams: map[string]string{},
	})

	assert.Equal(t, recorder.Code, http.StatusCreated, "Expected status code %d, got %d\n %v", http.StatusOK, recorder.Code, recorder.Body.String())

	deck := &models.Deck{}
	err := testutils.ParseRecorderBodyJson(*recorder, deck)

	assert.Nil(t, err, "Error parsing response body:\n %v", err)

	assert.Equal(t, deck.Name, deckName, "Expected deck name %s, got %s", deckName, deck.Name)

	assert.Conditionf(t, func() bool {
		err := uuid.Validate(deck.ID)

		if err != nil {
			return false
		}

		return true
	}, "invalid uuid in Deck.ID, expected a valid UUID, got %s", deck.ID)

	assert.Condition(t, func() bool {
		return len(deck.Cards) == 52
	}, "newly created decks should have 52 cards, got Deck.Cards as %v", deck.Cards)

}

func TestCreateInvalidDeck(t *testing.T) {
	// FIXME
}
