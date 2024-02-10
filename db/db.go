package db

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"togglDecks/common"
	"togglDecks/logging"
)

const (
	devDBPath     = "/dev.db.sqlite3"
	stagingDBPath = "/staging.db.sqlite3"
	prodDBPath    = "/prod.db.sqlite3"
)

func GetDBPath() string {
	switch {
	case common.CheckStage("dev"):
		return common.GetRootDir() + devDBPath
	case common.CheckStage("staging"):
		return common.GetRootDir() + stagingDBPath
	}

	return common.GetRootDir() + prodDBPath
}

func createDBFileIfNotExists(fileName string) error {
	_, err := os.Stat(fileName)

	if errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(fileName)

		if err != nil {
			return fmt.Errorf("failed to create DB file: %w", err)
		}

		defer file.Close()
	} else if err != nil {
		return fmt.Errorf("error checking if DB file exists: %w", err)
	}

	return nil
}

func LoadDb(fileName string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})

	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	if common.CheckStage("dev") || common.CheckStage("staging") {
		logging.Logger.Info("db debugging is active")
		return db.Debug()
	}

	return db
}

func GetDB() *gorm.DB {
	path := GetDBPath()

	logging.Logger.Info(fmt.Sprintf("using db: %s", path))

	err := createDBFileIfNotExists(path)

	if err != nil {
		logging.Logger.Fatal(err.Error())
	}

	return LoadDb(path)
}
