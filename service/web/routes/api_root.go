package routes

import (
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/gofiber/fiber/v2"
)

func init() {
	web.Register(func(svc *web.Service, app *fiber.App) {
		versionResponse, err := app.Config().JSONEncoder(fiber.Map{
			"application": constants.SumeruName,
			"version":     constants.SumeruVersion,
		})
		if err != nil {
			svc.Logger().DPanicf("failed to encode version response: %v", err)
			versionResponse = []byte("{}")
		}

		app.Get("/", func(ctx *fiber.Ctx) error {
			ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return ctx.Send(versionResponse)
		})
	})
}
