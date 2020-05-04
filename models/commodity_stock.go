package models

import (
	"time"
)

type CommodityStock struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	CommId         int       `xorm:"INT"`
	ColorId        int       `xorm:"INT"`
	SizeId         int       `xorm:"INT"`
	Stock          int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
