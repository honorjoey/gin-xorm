package models

import (
	"time"
)

type UserChange struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	Address        string    `xorm:"VARCHAR(256)"`
	Amount         float64   `xorm:"DOUBLE(18,4)"`
	Remark         string    `xorm:"VARCHAR(256)"`
	Fee            float64   `xorm:"DOUBLE(18,4)"`
	GetAmount      float64   `xorm:"DOUBLE(18,4)"`
	DrawHash       string    `xorm:"VARCHAR(256)"`
	ChangeSign     int       `xorm:"INT"`
	Status         int       `xorm:"INT"`
	CoinId         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Balance        float64   `xorm:"DOUBLE(18,4)"`
}
