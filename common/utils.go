package common

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
	"togglDecks/logging"
)

func CheckStage(stage string) bool {
	return strings.ToLower(os.Getenv("STAGE")) == strings.ToLower(stage)
}

func LoadEnvFile() {
	// this ignores the stage, .env is shared atm
	err := godotenv.Load(GetRootDir() + "/.env")

	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
}

func ParseCsvQueryString(queryString string) []string {
	return strings.Split(queryString, ",")
}
