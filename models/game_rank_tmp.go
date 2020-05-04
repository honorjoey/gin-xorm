package models

import (
	"time"
)

type GameRankTmp struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	MemberId   int       `xorm:"INT"`
	UserAccout string    `xorm:"VARCHAR(255)"`
	Asset      float64   `xorm:"DOUBLE(18,6)"`
	Score      int       `xorm:"INT"`
	CreateTime time.Time `xorm:"DATETIME"`
	GameId     int       `xorm:"INT"`
}
