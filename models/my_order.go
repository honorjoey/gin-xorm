package models

import (
	"time"
)

type MyOrder struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	OrderNum       string    `xorm:"VARCHAR(50)"`
	CommId         int       `xorm:"INT"`
	Details        string    `xorm:"VARCHAR(256)"`
	CommNum        int       `xorm:"INT"`
	Price          float64   `xorm:"DOUBLE(16,4)"`
	Freight        float64   `xorm:"DOUBLE(16,4)"`
	Express        string    `xorm:"VARCHAR(256)"`
	Address        string    `xorm:"VARCHAR(512)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Remarks        string    `xorm:"VARCHAR(255)"`
	ColorId        int       `xorm:"INT"`
	SizeId         int       `xorm:"INT"`
}
