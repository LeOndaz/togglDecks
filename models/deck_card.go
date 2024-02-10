package models

import (
	"fmt"
	"strings"
)

type DeckCard struct {
	ModelWithId

	Deck   Deck
	DeckID string

	// FIXME: should be here and this should only be a through-table
	// Card Card
	// CardID string

	// I don't really know how to play card games, I assumed Suit and Rank are a part of card games
	Order int
	Suit  string // FIXME: I advise against saving Enum strings in DB, but for the demo
	Rank  string // FIXME: Same here

	// FIXME: Code in API response?
}

func (card *DeckCard) Code() string {
	return strings.ToUpper(fmt.Sprintf("%c%c", card.Suit[0], card.Rank[0]))
}
