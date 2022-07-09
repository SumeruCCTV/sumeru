package r_auth

import (
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/middleware"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		log := svc.Logger().Named("auth")
		app.Post("/auth/register", middleware.Unauthorized(), func(ctx *fiber.Ctx) error {
			body, err := validateBody(ctx)
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
