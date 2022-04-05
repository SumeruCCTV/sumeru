package db

import (
	"github.com/SumeruCCTV/sumeru/service/database/models"
)

func (db *Database) RegisterAccount(username, password string) (*models.Account, error) {
	uuid := db.GenerateUuid()
	account := &models.Account{
		Uuid:         uuid,
		Username:     username,
		PasswordHash: password,
	}
	return account, db.Create(account).Error
}

func (db *Database) AccountByUsername(username string) (*models.Account, error) {
	var account *models.Account
	return account, db.Where("username = ?", username).First(&account).Error
}
