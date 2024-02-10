package routers

import (
	"github.com/gin-gonic/gin"
	"togglDecks/controllers/deck_controllers"
)

func SetupMainRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		_, _ = ctx.Writer.WriteString("PONG")
	})

	decks := router.Group("/decks")
	{
		decks.GET("", deck_controllers.GetAvailableDecks)
		decks.POST("", deck_controllers.NewDeck)
		decks.PUT("/:id/draw", deck_controllers.DrawCards)
		decks.GET("/:id", deck_controllers.OpenDeck)
	}

	return router
}
