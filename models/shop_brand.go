package models

import (
	"time"
)

type ShopBrand struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	Name       string `xorm:"VARCHAR(30)"`
	Img        string `xorm:"VARCHAR(256)"`
	GviName    string `xorm:"VARCHAR(30)"`
	GviDetails string `xorm:"VARCHAR(256)"`
	Phone      string `xorm:"VARCHAR(32)"`
	Details    string `xorm:"VARCHAR(1000)"`
	Status     int    `xorm:"comment('0 : app本身
            1：品牌介绍') INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
