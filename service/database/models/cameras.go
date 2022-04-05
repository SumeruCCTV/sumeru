package models

import "net"

type Camera struct {
	Uuid      string   `gorm:"primaryKey" json:"uuid"`
	Name      string   `gorm:"not null" json:"name"`
	IPAddress net.Addr `gorm:"not null" json:"addr"`
}
