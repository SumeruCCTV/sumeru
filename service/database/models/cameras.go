package models

type Camera struct {
	Uuid string `gorm:"primaryKey" json:"uuid"`

	// GORM currently dies when you set both "not null" and "uniqueIndex".
	OwnerUuid string `gorm:"uniqueIndex:idx_owner_uuid_name" json:"ownerUuid"`
	Name      string `gorm:"uniqueIndex:idx_owner_uuid_name" json:"name"`

	IPAddress   string            `gorm:"not null" json:"addr"`
	Port        int               `gorm:"not null" json:"port"`
	Type        CameraType        `gorm:"not null" json:"type"`
	Status      CameraStatus      `gorm:"not null" json:"status"`
	Credentials CameraCredentials `gorm:"not null;embedded;embeddedPrefix:camera_" json:"credentials"`

	CreatedAt int64 `json:"-"`
}

type CameraType int

const (
	CameraTypeONVIF CameraType = iota
	CameraTypeRTSP
)

type CameraStatus int

const (
	CameraStatusInvalid CameraStatus = iota
	CameraStatusDisconnected
	CameraStatusConnected
)

type CameraCredentials struct {
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}
