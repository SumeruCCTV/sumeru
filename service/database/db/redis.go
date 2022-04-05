package db

import (
	"context"
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/go-redis/redis/v8"
)

func rctx() context.Context {
	return context.Background()
}

func (db *Database) UuidFromToken(token string) (string, error) {
	return db.Redis.HGet(rctx(), constants.RedisTokenUuidKey, token).Result()
}

func (db *Database) SetTokenWithUuid(uuid string) (string, error) {
	ctx := rctx()
	token := db.GenerateToken()
	for err := db.Redis.HGet(ctx, constants.RedisTokenUuidKey, token).Err(); err != redis.Nil; {
		token = db.GenerateToken()
	}
	_, err := db.Redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(ctx, constants.RedisTokenUuidKey, token, uuid)
		pipe.Do(ctx, "EXPIREMEMBER", constants.RedisTokenUuidKey, token, constants.TokenExpiration)
		return nil
	})
	return token, err
}

func (db *Database) InvalidateToken(token string) error {
	return db.Redis.HDel(rctx(), constants.RedisTokenUuidKey, token).Err()
}
