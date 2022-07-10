package db

import (
	"github.com/SumeruCCTV/sumeru/service/database/models"
)

func (db *Database) AddCameraByUuid(camera *models.Camera) (*models.Camera, error) {
	camera.Uuid = db.GenerateUuid()
	camera.Status = models.CameraStatusInvalid
	return camera, db.Create(camera).Error
}

func (db *Database) UpdateCameraStatus(uuid string, status models.CameraStatus) error {
	// should we create the model only once?
	return db.Model(&models.Camera{}).
		Where("uuid = ?", uuid).
		Update("status", status).
		Error
}
