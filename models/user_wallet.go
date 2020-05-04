package models

import (
	"time"
)

type UserWallet struct {
	Id         int     `xorm:"not null pk autoincr INT"`
	Address    string  `xorm:"VARCHAR(256)"`
	Privatekey string  `xorm:"VARCHAR(256)"`
	UserId     int     `xorm:"index INT"`
	Amount     float64 `xorm:"DOUBLE(18,8)"`
	Status     int     `xorm:"comment('0：正常
            1：封禁') INT"`
	CoinId int `xorm:"comment('0：正常
            1：封禁') INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
