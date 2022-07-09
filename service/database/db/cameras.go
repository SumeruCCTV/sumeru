package db

import (
	"github.com/SumeruCCTV/sumeru/service/database/models"
)

func (db *Database) AddCameraByUuid(accountUuid, cameraName, cameraAddr string, cameraPort int, cameraType models.CameraType) (*models.Camera, error) {
	uuid := db.GenerateUuid()
	camera := &models.Camera{
		Uuid:      uuid,
		Name:      cameraName,
		OwnerUuid: accountUuid,
		IPAddress: cameraAddr,
		Port:      cameraPort,
		Type:      cameraType,
	}
	return camera, db.Create(camera).Error
}
