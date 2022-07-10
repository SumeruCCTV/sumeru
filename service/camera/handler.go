package camera

import (
	"github.com/SumeruCCTV/sumeru/service/database/models"
	"go.uber.org/zap"
)

func (svc *Service) accept(cam *models.Camera) {
	svc.log.Debugw("adding camera", zap.String("uuid", cam.Uuid))
	c := NewConnector(cam.Type, svc, &ConnectorData{
		uuid:        cam.Uuid,
		ipAddress:   cam.IPAddress,
		port:        cam.Port,
		credentials: cam.Credentials,
	})
	if c == nil {
		return // invalid camera type?
	}
	if err := c.TestConnection(); err != nil {
		svc.log.Debugw("failed to test connection for camera", zap.String("uuid", cam.Uuid), zap.Error(err))
		if cam.Status != models.CameraStatusInvalid {
			if err := svc.database.UpdateCameraStatus(cam.Uuid, models.CameraStatusInvalid); err != nil {
				svc.log.Warnw("failed to update camera status", zap.String("uuid", cam.Uuid), zap.Error(err))
			}
		}
		return
	}
	svc.cameras[cam.Uuid] = c
	if err := svc.database.UpdateCameraStatus(cam.Uuid, models.CameraStatusDisconnected); err != nil {
		svc.log.Warnw("failed to update camera status", zap.String("uuid", cam.Uuid), zap.Error(err))
		return
	}
}
