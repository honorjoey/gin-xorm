package models

import (
	"time"
)

type Message struct {
	Id        int    `xorm:"not null pk autoincr INT"`
	Name      string `xorm:"VARCHAR(256)"`
	Content   string `xorm:"VARCHAR(1000)"`
	NameEn    string `xorm:"VARCHAR(256)"`
	ContentEn string `xorm:"VARCHAR(1000)"`
	TopStatus int    `xorm:"comment('0：正常
            1：隐藏') INT"`
	Status int `xorm:"comment('0：正常
            1：隐藏') INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
