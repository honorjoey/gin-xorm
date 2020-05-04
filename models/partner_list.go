package models

import (
	"time"
)

type PartnerList struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	Price          float64   `xorm:"DOUBLE(16,4)"`
	Level          int       `xorm:"INT"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
