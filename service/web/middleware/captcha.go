package middleware

import (
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
)

func Captcha() fiber.Handler {
	if constants.IsDev {
		return func(ctx *fiber.Ctx) error {
			return ctx.Next() // disable captcha if in dev mode
		}
	}
	return func(ctx *fiber.Ctx) error {
		if !utils.HasCaptchaKey(ctx) {
			ctx.Status(fiber.StatusForbidden)
			return nil
		}
		return ctx.Next()
	}
}
