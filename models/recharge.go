package models

import (
	"time"
)

type Recharge struct {
	Id           int       `xorm:"not null pk autoincr INT"`
	UserId       int       `xorm:"not null comment('用户ID') INT"`
	CoinId       int       `xorm:"not null comment('币ID') INT"`
	RechargeDate time.Time `xorm:"DATETIME"`
	Amount       float64   `xorm:"default 0.0000 comment('充值数额') DOUBLE(16,4)"`
	BlockNum     int       `xorm:"comment('区块号') INT"`
	FromAddress  string    `xorm:"VARCHAR(255)"`
	ToAddress    string    `xorm:"VARCHAR(255)"`
	TxRemarks    string    `xorm:"VARCHAR(255)"`
	Coordinator  string    `xorm:"VARCHAR(255)"`
	ChangeSign   int       `xorm:"INT"`
	Balance      float64   `xorm:"DOUBLE(18,4)"`
	Status       int       `xorm:"default 0 INT"`
}
