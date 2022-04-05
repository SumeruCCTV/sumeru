package utils

import (
	"github.com/SumeruCCTV/sumeru/pkg/argon2id"
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/gofiber/fiber/v2"
)

func ValidUsername(username string, ctx *fiber.Ctx) error {
	if !IntBetween(len(username), constants.UsernameMinLength, constants.UsernameMaxLength) ||
		!constants.UsernameRegex.MatchString(username) {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidUsername
	}
	return nil
}

func ValidPassword(password string, ctx *fiber.Ctx) error {
	if !argon2id.ValidHash(password) {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidCredentials
	}
	return nil
}

func ValidBody(ctx *fiber.Ctx, tests ...string) error {
	for _, t := range tests {
		if StringBlank(t) {
			ctx.Status(fiber.StatusBadRequest)
			return errors.InvalidBody
		}
	}
	return nil
}

var byteHeaderKey = []byte(constants.HeaderCaptchaKey)

func HasCaptchaKey(ctx *fiber.Ctx) bool {
	return len(ctx.Request().Header.PeekBytes(byteHeaderKey)) != 0
}
