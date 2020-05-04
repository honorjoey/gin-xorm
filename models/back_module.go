package models

import (
	"time"
)

type BackModule struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Name           string    `xorm:"VARCHAR(125)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
