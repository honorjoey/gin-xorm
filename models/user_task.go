package models

import (
	"time"
)

type UserTask struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	TaskId         int       `xorm:"index INT"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Amount         float64   `xorm:"DOUBLE(18,8)"`
}
