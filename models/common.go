package models

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ModelWithId struct {
	ID string
}

func (instance *ModelWithId) BeforeCreate(tx *gorm.DB) error {
	// FIXME: I don't like using UUID in IDs because they are not sortable, however, UUIDv7 is, there are also snowflakes

	id, err := uuid.NewV7()

	if err != nil {
		return errors.New(fmt.Sprintf("failed to create UUID for %s", tx.Model(instance).Name()))
	}

	instance.ID = id.String()
	return nil
}
