package r_auth

import (
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
)

type requestBody struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func validateBody(ctx *fiber.Ctx) (body requestBody, err error) {
	if err = ctx.BodyParser(&body); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return body, errors.InvalidBody
	}
	if err = utils.ValidBody(ctx, body.Username, body.PasswordHash); err != nil {
		return
	}
	if err = utils.ValidUsername(body.Username, ctx); err != nil {
		return
	}
	if err = utils.ValidPassword(body.PasswordHash, ctx); err != nil {
		return
	}
	return
}
