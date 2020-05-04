package models

import (
	"time"
)

type Block struct {
	Id         int       `xorm:"not null pk autoincr INT"`
	CoinId     int       `xorm:"not null comment('币ID') INT"`
	UpdateDate time.Time `xorm:"comment('更新时间') DATETIME"`
	BlockNum   int       `xorm:"comment('区块号') INT"`
	SyncId     int       `xorm:"INT"`
}
