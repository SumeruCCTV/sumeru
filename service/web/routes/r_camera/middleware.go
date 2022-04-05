package r_camera

import (
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/middleware"
	"github.com/gofiber/fiber/v2"
)

func init() {
	web.RegisterMiddleware(func(app *fiber.App) {
		app.Use("/camera", middleware.Authorized())
	})
}
