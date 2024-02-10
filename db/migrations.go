package db

import (
	"togglDecks/logging"
	"togglDecks/models"
)

var modelToMigrate = []any{
	&models.Deck{},
	&models.Card{},
	&models.DeckCard{},
}

func MigrateDb() {
	err := GetDB().AutoMigrate(modelToMigrate...)

	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
}
