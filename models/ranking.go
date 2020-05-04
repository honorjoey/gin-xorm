package models

import (
	"time"
)

type Ranking struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Ranking        int       `xorm:"INT"`
	UserId         int       `xorm:"INT"`
	Nickname       string    `xorm:"VARCHAR(30)"`
	Identitys      string    `xorm:"VARCHAR(30)"`
	Amount         float64   `xorm:"DOUBLE(16,4)"`
	Phone          string    `xorm:"VARCHAR(30)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
