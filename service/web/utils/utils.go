package utils

import (
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"strings"
)

const emptyString = ""

var byteCookieKey = []byte(constants.CookieTokenKey)

func UuidFromCtx(ctx *fiber.Ctx, svc *web.Service) (string, string, error) {
	token := ctx.Cookies(constants.CookieTokenKey)
	uuid, err := svc.DB().UuidFromToken(token)
	if err != nil || uuid == emptyString {
		if err != redis.Nil {
			// only log if it's an error that's useful to know
			svc.Logger().With("token", token).
				Errorf("failed to get uuid from token: %v", err)
		}
		ctx.Status(fiber.StatusUnauthorized)
		return emptyString, emptyString, errors.InvalidToken
	}
	return uuid, token, nil
}

func IsValidToken(ctx *fiber.Ctx) bool {
	token := ctx.Request().Header.CookieBytes(byteCookieKey)
	return len(token) == constants.TokenLength && constants.TokenRegex.Match(token)
}

func IntBetween(i, min, max int) bool {
	return i >= min && i <= max
}

func StringBlank(s string) bool {
	return s == "" || strings.TrimSpace(s) == ""
}
