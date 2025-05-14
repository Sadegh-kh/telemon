package idgenerator

import "github.com/google/uuid"

type UUIDGenerator struct {
}

func (g *UUIDGenerator) GenerateID() string {
	return uuid.New().String()
}
