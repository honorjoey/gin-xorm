package models

import (
	"time"
)

type TasksIncome struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	TaskId         int       `xorm:"INT"`
	Income         float64   `xorm:"DOUBLE(16,4)"`
	CoinId         int       `xorm:"INT"`
	Types          int       `xorm:"INT"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Balance        float64   `xorm:"DOUBLE(18,4)"`
}
