package r_camera

import (
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
)

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		log := svc.Logger().Named("camera")
		app.Post("/camera", func(ctx *fiber.Ctx) error {
			uuid, _, err := utils.UuidFromCtx(ctx, svc)
			if err != nil {
				return err
			}
			log.Debugf("hello user with uuid: %s", uuid)
			return nil
		})
	})
}
