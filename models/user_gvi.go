package models

import (
	"time"
)

type UserGvi struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"index INT"`
	Amount         float64   `xorm:"DOUBLE(16,4)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LsatUpdateTime time.Time `xorm:"DATETIME"`
}
