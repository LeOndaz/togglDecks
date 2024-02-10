package tests

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"testing"
	"togglDecks/db"
)

func TestCreateDbFileIfNotExists(t *testing.T) {
	const fakeDbName = "fake.db.sqlite3"
	db.LoadDb(fakeDbName)

	defer os.Remove(fakeDbName)

	_, err := os.Stat(fakeDbName)

	assert.Nil(t, err, "expected file=%s to be created", fakeDbName)
}

func TestLoadDb(t *testing.T) {
	const fakeDbName = "fake.db.sqlite3"

	conn := db.LoadDb(fakeDbName)
	defer os.Remove(fakeDbName)

	assert.IsType(t, &gorm.DB{}, conn, "expected conn to be of type gorm.DB")
}
