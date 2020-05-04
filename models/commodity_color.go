package models

import (
	"time"
)

type CommodityColor struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	CommId         int       `xorm:"INT"`
	Name           string    `xorm:"VARCHAR(50)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
