package main

import (
	"github.com/google/uuid"
)

func CreateUUID() (uuid.UUID, error) {
	return uuid.NewV7()
}
