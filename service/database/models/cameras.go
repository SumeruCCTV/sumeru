package models

type Camera struct {
	Uuid      string `gorm:"primaryKey,unique_index:idx_uuid2name" json:"uuid"`
	OwnerUuid string `gorm:"not null" json:"ownerUuid"`
	Name      string `gorm:"not null,unique_index:idx_uuid2name" json:"name"`
	IPAddress string `gorm:"not null" json:"addr"`
	CreatedAt int64  `json:"-"`
}
