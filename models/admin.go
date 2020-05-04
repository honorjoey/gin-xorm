package models

import (
	"time"
)

type Admin struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Username       string    `xorm:"VARCHAR(256)"`
	Phone          string    `xorm:"VARCHAR(256)"`
	Password       string    `xorm:"VARCHAR(256)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	RoolId         int       `xorm:"INT"`
}
