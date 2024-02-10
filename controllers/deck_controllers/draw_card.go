package deck_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"togglDecks/common"
	"togglDecks/db"
	"togglDecks/logging"
	"togglDecks/services/decks"
)

func DrawCards(ctx *gin.Context) {
	count := 1
	countQuery := ctx.Query("count")

	if countQuery != "" {
		parsedCount, err := strconv.ParseUint(countQuery, 10, 64)

		if err != nil {
			logging.Logger.Error(err.Error())

			ctx.JSON(http.StatusBadRequest, common.Error{
				Code:    common.InvalidQueryParameter,
				Message: fmt.Sprintf("invalid query parameter count=%s", countQuery),
			})
			return
		}

		count = int(parsedCount)
	}

	deckId := ctx.Param("id")
	deck, err := decks.GetDeckById(deckId)

	if err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(http.StatusNotFound, common.Error{
			Code:    common.DeckNotFound,
			Message: fmt.Sprintf("deck with id=%s not found", deckId),
		})
		return
	}

	cards, err := deck.Draw(count)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:    common.DeckFailedToDraw,
			Message: err.Error(),
			Metadata: gin.H{
				"deck": deck,
			},
		})
		return
	}

	if err := db.GetDB().Delete(cards).Error; err != nil {
		logging.Logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:     common.DeckFailedToDraw,
			Message:  fmt.Sprintf("couldn't draw from deck=%s", deck.ID),
			Metadata: nil,
		})
	}

	if err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:    common.DeckFailedToDraw,
			Message: fmt.Sprintf("couldn't draw cards from deck=%s", deckId),
			Metadata: gin.H{
				"deck": deck,
			},
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"cards":     cards,
		"countLeft": len(deck.Cards),
	})
}
