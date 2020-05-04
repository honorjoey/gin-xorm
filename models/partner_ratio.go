package models

import (
	"time"
)

type PartnerRatio struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Name           string    `xorm:"VARCHAR(56)"`
	Price          float64   `xorm:"DOUBLE(16,4)"`
	Ratio          float64   `xorm:"DOUBLE(16,4)"`
	NewRatio       float64   `xorm:"DOUBLE(16,4)"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
