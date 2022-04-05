package r_auth

import (
	errs "errors"
	"github.com/SumeruCCTV/sumeru/pkg/argon2id"
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/middleware"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type loginResponse struct {
	Token string `json:"token"`
}

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		log := svc.Logger().Named("auth")
		app.Post("/auth/login", middleware.Unauthorized(), func(ctx *fiber.Ctx) error {
			body, err := authValidateBody(ctx)
			if err != nil {
				return err
			}
			account, err := svc.DB().AccountByUsername(body.Username)
			if err != nil {
				if !errs.Is(err, gorm.ErrRecordNotFound) || account == nil {
					log.With(zap.String("username", body.Username)).
						Errorf("error finding account in db: %v", err)
					return errors.ErrorLoggingIn
				}
				ctx.Status(fiber.StatusBadRequest)
				return errors.InvalidCredentials
			}
			match, err := argon2id.CompareHashes(body.PasswordHash, account.PasswordHash)
			if err != nil {
				log.With(
					zap.String("username", body.Username),
					zap.String("uuid", account.Uuid),
				).Errorf("error checking hash: %v", err)
				return errors.ErrorLoggingIn
			}
			if !match {
				ctx.Status(fiber.StatusBadRequest)
				return errors.InvalidCredentials
			}
			token, err := svc.DB().SetTokenWithUuid(account.Uuid)
			if err != nil {
				log.With(
					zap.String("username", body.Username),
					zap.String("uuid", account.Uuid),
				).Errorf("error setting token: %v", err)
				return errors.ErrorLoggingIn
			}
			return ctx.JSON(loginResponse{token})
		})
	})
}
