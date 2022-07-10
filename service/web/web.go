package web

import (
	"fmt"
	"github.com/SumeruCCTV/sumeru/pkg/config"
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/pkg/json"
	"github.com/SumeruCCTV/sumeru/pkg/svcstat"
	"github.com/SumeruCCTV/sumeru/pkg/utils"
	"github.com/SumeruCCTV/sumeru/service/camera"
	"github.com/SumeruCCTV/sumeru/service/database"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
)

type Service struct {
	cfg      *config.Config
	log      *utils.Logger
	database *database.Service
	camera   *camera.Service

	app     *fiber.App
	running bool
}

func (Service) Name() string {
	return "web"
}

func (svc *Service) Start() error {
	svc.app = fiber.New(fiber.Config{
		AppName:               constants.SumeruName,
		DisableStartupMessage: true,
		DisableDefaultDate:    true,
		// Use faster json serialization and de-serialization
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		// Custom error handler to handle errors.WebError
		ErrorHandler: errorHandler,
		BodyLimit:    1024 * 1024, // 1 MB
	})
	svc.app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	svc.app.Use(helmet.New())
	svc.app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,PUT,POST,DELETE,OPTIONS",
		AllowCredentials: true,
	}))
	svc.app.Use(requestid.New())
	// TODO: do we want to compress even if the request is invalid?
	svc.app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	// TODO: use our own logger
	svc.app.Use(logger.New(logger.Config{Format: "[${time}] ${locals:requestid} ${status} - ${latency} ${method} ${path}\n"}))
	registerMiddleware(svc.app)
	registerRoutes(svc, svc.app)
	svc.app.Use(notFoundHandler)

	go func() {
		// todo: return error
		port := svc.cfg.Web.Port
		svc.log.Infof("Starting web server on port %d", port)
		svc.running = true
		_ = svc.app.Listen(fmt.Sprintf(":%d", port))
		svc.running = false
	}()
	return nil
}

func (svc *Service) Stop() error {
	return svc.app.Shutdown()
}

func (svc *Service) Status() svcstat.Status {
	if svc.app != nil && svc.database != nil && svc.running {
		return svcstat.StatusHealthy
	}
	return svcstat.StatusUnhealthy
}

func (svc *Service) Logger() *utils.Logger {
	return svc.log
}

func (svc *Service) DB() *database.Service {
	return svc.database
}

func (svc *Service) CameraSvc() *camera.Service {
	return svc.camera
}

var errorHandler = func(ctx *fiber.Ctx, _e error) error {
	if ctx.Response().StatusCode() == fiber.StatusOK {
		ctx.Status(fiber.StatusInternalServerError)
	}

	if err, ok := _e.(errors.WebError); ok {
		return ctx.Send(err)
	}
	return ctx.JSON(json.Error(_e.Error()))
}

var notFoundHandler = func(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusNotFound)
	return errors.NotFound
}
