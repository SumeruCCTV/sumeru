package db

import (
	"github.com/SumeruCCTV/sumeru/service/database/models"
)

func (db *Database) AddCameraByUuid(accountUuid, cameraName, cameraAddr string) (*models.Camera, error) {
	uuid := db.GenerateUuid()
	camera := &models.Camera{
		Uuid:      uuid,
		OwnerUuid: accountUuid,
		Name:      cameraName,
		IPAddress: cameraAddr,
	}
	return camera, db.Create(camera).Error
}
