package models

import (
	"time"
)

type UserVote struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	FranId         int       `xorm:"INT"`
	GviValue       float64   `xorm:"DOUBLE(16,4)"`
	Ratio          int       `xorm:"INT"`
	Value          float64   `xorm:"DOUBLE(16,4)"`
	Types          int       `xorm:"INT"`
	Status         int       `xorm:"comment('0：正在释放  1：释放完成') INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
