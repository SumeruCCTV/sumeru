package db

import (
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/go-redis/redis/v8"
	nanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
	Redis *redis.Client
}

func (Database) GenerateUuid() string {
	uuid, _ := nanoid.Generate(constants.UuidTokenAlphabet, constants.UuidLength)
	return uuid
}

func (Database) GenerateToken() string {
	token, _ := nanoid.Generate(constants.UuidTokenAlphabet, constants.TokenLength)
	return token
}
