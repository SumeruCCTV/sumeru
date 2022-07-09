package utils

import (
	"github.com/SumeruCCTV/sumeru/pkg/argon2id"
	"github.com/SumeruCCTV/sumeru/pkg/constants"
	"github.com/SumeruCCTV/sumeru/pkg/errors"
	"github.com/SumeruCCTV/sumeru/service/database/models"
	"github.com/gofiber/fiber/v2"
	"net"
)

func ValidUsername(username string, ctx *fiber.Ctx) error {
	if !IntBetween(len(username), constants.UsernameMinLength, constants.UsernameMaxLength) ||
		!constants.UsernameRegex.MatchString(username) {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidUsername
	}
	return nil
}

func ValidPassword(password string, ctx *fiber.Ctx) error {
	if !argon2id.ValidHash(password) {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidCredentials
	}
	return nil
}

func ValidCameraName(name string, ctx *fiber.Ctx) error {
	if !IntBetween(len(name), constants.CameraNameMinLength, constants.CameraNameMaxLength) ||
		!constants.CameraNameRegex.MatchString(name) {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidCameraName
	}
	return nil
}

func ValidCameraAddr(addr string, ctx *fiber.Ctx) error {
	if net.ParseIP(addr) == nil {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidCameraAddr
	}
	return nil
}

func ValidCameraPort(port int, ctx *fiber.Ctx) error {
	// 80 sounds like a sensible minimum port number
	if port < 80 || port > 65535 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidCameraPort
	}
	return nil
}

func ValidCameraType(cameraType models.CameraType, ctx *fiber.Ctx) error {
	if cameraType < models.CameraTypeUnknown || cameraType > models.CameraTypeRTSP {
		ctx.Status(fiber.StatusBadRequest)
		return errors.InvalidCameraType
	}
	return nil
}

func ValidBody(ctx *fiber.Ctx, tests ...string) error {
	for _, t := range tests {
		if StringBlank(t) {
			ctx.Status(fiber.StatusBadRequest)
			return errors.InvalidBody
		}
	}
	return nil
}

var byteHeaderKey = []byte(constants.HeaderCaptchaKey)

func HasCaptchaKey(ctx *fiber.Ctx) bool {
	return len(ctx.Request().Header.PeekBytes(byteHeaderKey)) != 0
}
