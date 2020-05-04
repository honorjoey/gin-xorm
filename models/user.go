package models

import (
	"time"
)

type User struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Areacode       string    `xorm:"VARCHAR(10)"`
	Telphone       string    `xorm:"VARCHAR(20)"`
	Password       string    `xorm:"VARCHAR(40)"`
	Tradpassword   string    `xorm:"comment('交易密码') VARCHAR(255)"`
	Nickname       string    `xorm:"VARCHAR(30)"`
	Avatar         string    `xorm:"VARCHAR(255)"`
	InvestCode     string    `xorm:"VARCHAR(20)"`
	Star           int       `xorm:"INT"`
	Status         int       `xorm:"default 0 comment('0 正常 1被禁用') INT"`
	LoginNum       int       `xorm:"default 0 comment('0 正常 1被禁用') INT"`
	LastLoginTime  time.Time `xorm:"DATETIME"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Inviter        int       `xorm:"INT"`
}
