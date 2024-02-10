package models

import (
	"errors"
	"fmt"
	"math/rand"
	"togglDecks/common"
)

type Deck struct {
	ModelWithId

	Name string

	Cards []DeckCard

	Shuffled bool `gorm:"default:false;"`
}

// Reset I didn't implement this as a hook because I assumed that we may have empty decks at some point
func (deck *Deck) Reset(requiredCards []string) error {
	// I got stuck in this function because I thought arrays are pointers to 1st element in GoLang (Since I wrote lots of C after GoLang)

	if deck == nil {
		return common.MustBeCalledOnNonNil("Deck.Reset()")
	}

	if deck.ID == "" {
		return errors.New("deck has no ID, must call Save() first")
	}

	if deck.Shuffled {
		// FIXME: should handle a case where this is called on a shuffled deck
	}

	var cardsCount int

	if requiredCards != nil {
		cardsCount = len(requiredCards)
	} else {
		cardsCount = len(common.Suits) * len(common.Ranks)
	}

	println("#######")
	fmt.Printf("%v\n", requiredCards)

	println(len(requiredCards))
	println(cardsCount)

	cards := make([]DeckCard, cardsCount)

	for i, code := range requiredCards {
		rank, err := common.GetRankFullName(code[0])

		if err != nil {
			return errors.New(fmt.Sprintf("invalid rank=%c provided for card at index=%d", code[0], i))
		}

		suit, err := common.GetSuitFullName(code[1])

		if err != nil {
			return errors.New(fmt.Sprintf("invalid suit=%c provided for card at index=%d", code[1], i))
		}

		card := DeckCard{
			DeckID: deck.ID,
			Suit:   suit,
			Rank:   rank,
			Order:  i, // order of creation
		}
		cards[i] = card
	}

	deck.Cards = cards

	return nil
}

func (deck *Deck) Shuffle() error {
	// shuffle the deck cards, using .Cards

	if deck == nil {
		return common.MustBeCalledOnNonNil("Deck.Shuffle()")
	}

	for i := 0; i < len(deck.Cards); i++ {
		j := rand.Intn(i + 1)
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
		deck.Cards[i].Order = j
		deck.Cards[j].Order = i
	}

	deck.Shuffled = true

	return nil
}

func (deck *Deck) Draw(count int) ([]DeckCard, error) {
	if deck == nil {
		return nil, common.MustBeCalledOnNonNil("Deck.Draw()")
	}

	if deck.Cards == nil {
		return nil, errors.New("called Deck.Draw() on empty deck")
	}

	if count < 0 {
		return nil, errors.New("really")
	}

	if count == 0 {
		return nil, errors.New("must draw at least one card")
	}

	if count > len(deck.Cards) {
		return nil, errors.New("cannot draw more cards than are in the deck")
	}

	// draw = delete, for now
	lastIdx := len(deck.Cards)
	cardsToDraw := deck.Cards[lastIdx-count:]

	// not querying the deck again, so updating it here
	deck.Cards = deck.Cards[:lastIdx-count]
	return cardsToDraw, nil
}
