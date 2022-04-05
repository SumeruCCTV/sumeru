package middleware

import (
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
)

func Authorized(callback ...fiber.Handler) fiber.Handler {
	var cb fiber.Handler = nil
	if len(callback) > 0 {
		cb = callback[0]
	} else {
		cb = func(ctx *fiber.Ctx) error {
			ctx.Status(fiber.StatusUnauthorized)
			return nil
		}
	}
	return func(ctx *fiber.Ctx) error {
		if !utils.IsValidToken(ctx) {
			// unauthorized, invoke callback
			return cb(ctx)
		}
		return ctx.Next()
	}
}
