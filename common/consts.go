package common

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"togglDecks/logging"
)

// Suits are ordered
var Suits = [4]string{"Spades", "Diamonds", "Clubs", "Hearts"}
var Ranks = [13]string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

const CardsCount = len(Suits) * len(Ranks)

func getCardCodes() [CardsCount]string {
	codes := [CardsCount]string{}

	for i, suit := range Suits {
		for j, rank := range Ranks {
			idx := (i * len(Ranks)) + j
			codes[idx] = strings.ToUpper(fmt.Sprintf("%c%c", rank[0], suit[0]))
		}
	}

	return codes
}

var AvailableCardCodes = getCardCodes()

// FIXME: I tried for two hours to get it dynamically using runtime.Caller(1)
// it fails so I just ignored it for now for the demo

func GetRootDir() string {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		logging.Logger.Fatal("Failed to get root directory")
	}

	logging.Logger.Info(fmt.Sprintf("Root directory: %s", filename))
	return path.Dir(path.Dir(filename))
}
