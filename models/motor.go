package models

import (
	"time"
)

type Motor struct {
	ID         int64     `gorm:"primaryKey" json:"id"`
	BrandMotor string    `gorm:"type:varchar(300)" json:"brandMotor"`
	NamaMotor  string    `gorm:"type:varchar(300)" json:"namaMotor"`
	TipeMotor  string    `gorm:"type:varchar(300)" json:"tipeMotor"`
	PlatNomor  string    `gorm:"type:varchar(300)" json:"platMotor"`
	TahunPlat  time.Time `json:"tahunPlat"`
}

type Kaki_Kaki struct {
	ID      int64 `gorm:"primaryKey" json:"id"`
	MotorID int64 ``
}
