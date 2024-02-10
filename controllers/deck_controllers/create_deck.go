package deck_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"togglDecks/common"
	"togglDecks/db"
	"togglDecks/logging"
	"togglDecks/models"
	"togglDecks/services/decks"
)

func NewDeck(ctx *gin.Context) {
	deck := &models.Deck{}

	if err := ctx.ShouldBindJSON(deck); err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(http.StatusUnprocessableEntity, common.Error{
			Code:    common.ValidationError,
			Message: "invalid deck body provided",
		})
		return
	}

	// one query here
	validationErr := decks.InitEmptyDeck(deck)

	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:    common.UnknownError,
			Message: validationErr.Error(),
		})
		return
	}

	cardsQuery, ok := ctx.GetQuery("cards")
	cards := common.ParseCsvQueryString(cardsQuery)

	if ok == false {
		cards = common.AvailableCardCodes[:]
	}

	if err := deck.Reset(cards); err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(http.StatusBadRequest, common.Error{
			Code:    common.UnknownError,
			Message: err.Error(),
		})
		return
	}

	shuffle := ctx.Query("shuffle")

	if shuffle != "" && shuffle != "false" && shuffle != "0" && shuffle != "no" { // this logic should be in a utility
		_ = deck.Shuffle()
	}

	// one query happens here
	if err := db.GetDB().Save(&deck.Cards).Error; err != nil {
		logging.Logger.Error(err.Error())

		ctx.JSON(400, common.Error{
			Code:    common.UnknownError,
			Message: "couldn't save deck cards, please try again",
		})
		return
	}

	ctx.JSON(http.StatusCreated, deck)
}
