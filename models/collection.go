package models

import (
	"time"
)

type Collection struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	CoinId         int       `xorm:"not null comment('币ID') INT"`
	FromAddress    string    `xorm:"comment('用户地址') VARCHAR(255)"`
	Amount         float64   `xorm:"default 0.0000 comment('数额') DOUBLE(16,4)"`
	CollectionDate time.Time `xorm:"comment('时间') DATETIME"`
	TxHash         string    `xorm:"comment('交易哈希') VARCHAR(255)"`
	TxRemarks      string    `xorm:"default '' comment('交易备注') VARCHAR(255)"`
}
