package models

type Account struct {
	Uuid         string `gorm:"primaryKey" json:"uuid"`
	Username     string `gorm:"not null;unique" json:"username"`
	PasswordHash string `gorm:"not null" json:"-"`
	CreatedAt    int64  `json:"-"`
}
