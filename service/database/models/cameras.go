package models

type Camera struct {
	Uuid      string `gorm:"primaryKey,unique_index:idx_uuid2name" json:"uuid"`
	Name      string `gorm:"not null,unique_index:idx_uuid2name" json:"name"`
	OwnerUuid string `gorm:"not null" json:"ownerUuid"`

	IPAddress string     `gorm:"not null" json:"addr"`
	Port      int        `gorm:"not null" json:"port"`
	Type      CameraType `gorm:"not null" json:"type"`

	CreatedAt int64 `json:"-"`
}

type CameraType int

const (
	CameraTypeUnknown CameraType = iota
	CameraTypeONVIF
	CameraTypeRTSP
)
