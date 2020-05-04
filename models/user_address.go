package models

import (
	"time"
)

type UserAddress struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	Consignee      string    `xorm:"VARCHAR(50)"`
	Phone          string    `xorm:"VARCHAR(50)"`
	Provinces      string    `xorm:"VARCHAR(50)"`
	Cities         string    `xorm:"VARCHAR(50)"`
	Counties       string    `xorm:"VARCHAR(50)"`
	Address        string    `xorm:"VARCHAR(512)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
