package tests

import (
	"testing"
	"togglDecks/models"
	"togglDecks/services/decks"
)

func TestValidateValidDeck(t *testing.T) {
	deck := &models.Deck{
		Name:     "Gojo Satoru",
		Shuffled: false,
	}
	err := decks.ValidateDeck(deck)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestValidateInvalidDeck(t *testing.T) {
	deck := &models.Deck{
		Name:     "",
		Shuffled: false,
	}

	err := decks.ValidateDeck(deck)

	if err == nil {
		t.Errorf("Expected to fail, got %v", err)
	}
}
