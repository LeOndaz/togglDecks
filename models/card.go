package models

// Card FIXME: just to make sure we can have cards unrelated to Decks
// Also this can help if we "draw" cards from decks, so cards get removed (forever)
// If we had this working, cards would stay outside decks without issues
type Card struct {
	ModelWithId

	Deck   Deck
	DeckID string
}
