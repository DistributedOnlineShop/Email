package util

import (
	"log"

	"github.com/google/uuid"
)

func CreateUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		log.Fatal(err)
	}
	return id
}
