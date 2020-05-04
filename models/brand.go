package models

import (
	"time"
)

type Brand struct {
	Id         int    `xorm:"not null pk autoincr INT"`
	Name       string `xorm:"VARCHAR(30)"`
	Img        string `xorm:"VARCHAR(256)"`
	GviName    string `xorm:"VARCHAR(30)"`
	GviDetails string `xorm:"VARCHAR(256)"`
	Phone      string `xorm:"VARCHAR(32)"`
	Details    string `xorm:"VARCHAR(1000)"`
	Status     int    `xorm:"comment('0 : app本身
            1：品牌介绍') INT"`
	Quota int `xorm:"comment('0 : app本身
            1：品牌介绍') INT"`
	ExIncome       string    `xorm:"VARCHAR(256)"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
