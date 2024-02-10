package deck_controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"togglDecks/common"
	"togglDecks/logging"
	"togglDecks/services/decks"
)

func OpenDeck(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:     common.InvalidPathParameter,
			Message:  "didn't provide an ID",
			Metadata: nil,
		})
	}

	deck, err := decks.GetDeckById(id)

	if err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(400, common.Error{
			Code:    common.CardNotFound,
			Message: fmt.Sprintf("deck with id=%s was not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, deck)
}
