package models

import (
	"time"
)

type BackRoot struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	Name           string    `xorm:"VARCHAR(125)"`
	Tomarks        string    `xorm:"VARCHAR(125)"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
