package r_auth

import (
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/middleware"
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type authBody struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
}

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		log := svc.Logger().Named("auth")
		app.Post("/auth/register", middleware.Unauthorized(), func(ctx *fiber.Ctx) error {
			body, err := authValidateBody(ctx)
			if err != nil {
				return err
			}
			account, err := svc.DB().RegisterAccount(body.Username, body.PasswordHash)
			if err != nil {
				if errors.IsPgErr(err, errors.PgErrDuplicateEntry) {
					ctx.Status(fiber.StatusConflict)
					return nil
				}
				uuid := "unknown"
				if account != nil {
					uuid = account.Uuid
				}
				log.With(
					zap.String("username", body.Username),
					zap.String("uuid", uuid),
				).Errorf("error registering account in db: %v", err)
				return errors.ErrorRegisteringAccount
			}
			ctx.Status(fiber.StatusCreated)
			return nil
		})
	})
}

func authValidateBody(ctx *fiber.Ctx) (body authBody, err error) {
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
