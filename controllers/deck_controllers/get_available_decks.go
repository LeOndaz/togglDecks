package deck_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"togglDecks/common"
	"togglDecks/db"
	"togglDecks/logging"
	"togglDecks/models"
)

// FIXME: this was not requested, but I didn't like getting IDs from the DB (even though the DB is built-in the IDE)
// not covered by tests

func GetAvailableDecks(ctx *gin.Context) {
	// two queries

	var decks []models.Deck

	if err := db.GetDB().Where(&models.Deck{}).Preload("Cards").Limit(100).Find(&decks).Error; err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:    common.UnknownError,
			Message: "error while getting list of decks",
		})
	}

	ctx.JSON(http.StatusOK, decks)
}
