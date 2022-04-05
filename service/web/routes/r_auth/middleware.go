package r_auth

import (
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

var _authLimiter = limiter.New(limiter.Config{
	Next: func(*fiber.Ctx) bool {
		return constants.IsDev
	},
	Max:        1,
	Expiration: 10 * time.Minute,
	LimitReached: func(c *fiber.Ctx) error {
		c.Status(fiber.StatusTooManyRequests)
		return errors.TooManyRequests
	},
	// TODO: use redis storage
	// Storage: redisStorage{}
})

func init() {
	web.RegisterMiddleware(func(app *fiber.App) {
		app.Route("/auth", func(r fiber.Router) {
			r.Use("/register", _authLimiter, middleware.Captcha())
			r.Use("/login", _authLimiter, middleware.Captcha())
		})
	})
}
