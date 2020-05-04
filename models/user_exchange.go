package models

import (
	"time"
)

type UserExchange struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	UserId         int       `xorm:"INT"`
	CoinId         int       `xorm:"INT"`
	TocoinId       int       `xorm:"INT"`
	CoinRatio      float64   `xorm:"DOUBLE(16,4)"`
	TocoinRatio    float64   `xorm:"DOUBLE(16,4)"`
	Amount         float64   `xorm:"DOUBLE(16,4)"`
	Ratio          string    `xorm:"VARCHAR(16)"`
	Fee            float64   `xorm:"DOUBLE(16,4)"`
	GetAmount      float64   `xorm:"DOUBLE(16,4)"`
	Status         int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Balance        float64   `xorm:"DOUBLE(18,4)"`
}
