package r_auth

import (
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/middleware"
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		log := svc.Logger().Named("auth")
		app.Post("/auth/logout", middleware.Authorized(), func(ctx *fiber.Ctx) error {
			_, token, err := utils.UuidFromCtx(ctx, svc)
			if err != nil {
				return err
			}
			err = svc.DB().InvalidateToken(token)
			if err != nil {
				log.Errorf("failed to invalidate token: %v", err)
				return errors.New(errors.ErrorLoggingOut)
			}
			return nil
		})
	})
}
