package r_camera

import (
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/web"
	"github.com/SumeruCCTV/sumeru/service/web/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func init() {
	type requestBody struct {
		Name string `json:"cameraName"`
		Addr string `json:"cameraAddr"`
	}

	type responseBody struct {
		Uuid string `json:"uuid"`
	}

	validateBody := func(ctx *fiber.Ctx) (body requestBody, err error) {
		if err = ctx.BodyParser(&body); err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return body, errors.InvalidBody
		}
		if err = utils.ValidBody(ctx, body.Name, body.Addr); err != nil {
			return
		}
		if err = utils.ValidCameraName(body.Name, ctx); err != nil {
			return
		}
		if err = utils.ValidCameraAddr(body.Addr, ctx); err != nil {
			return
		}
		return
	}

	web.Register(func(svc *web.Service, app *fiber.App) {
		log := svc.Logger().Named("camera")
		app.Post("/camera", func(ctx *fiber.Ctx) error {
			uuid, _, err := utils.UuidFromCtx(ctx, svc)
			if err != nil {
				return err
			}
			body, err := validateBody(ctx)
			if err != nil {
				return err
			}
			camera, err := svc.DB().AddCameraByUuid(uuid, body.Name, body.Addr)
			if err != nil {
				if errors.IsPgErr(err, errors.PgErrDuplicateEntry) {
					ctx.Status(fiber.StatusConflict)
					return nil
				}
				log.With(
					zap.String("name", body.Name),
					zap.String("uuid", uuid),
				).Errorf("error adding camera to db: %v", err)
				return errors.ErrorAddingCamera
			}
			ctx.Status(fiber.StatusCreated)
			return ctx.JSON(responseBody{Uuid: camera.Uuid})
		})
	})
}
