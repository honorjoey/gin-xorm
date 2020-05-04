package models

import (
	"time"
)

type RoolToken struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	RoolId         int       `xorm:"INT"`
	ModId          int       `xorm:"INT"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
