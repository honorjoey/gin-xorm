package models

type Chain struct {
	Id      int    `xorm:"not null pk autoincr INT"`
	CoinId  int    `xorm:"not null comment('币ID') INT"`
	NodeNet string `xorm:"default '' comment('节点地址') VARCHAR(255)"`
	Remarks string `xorm:"default '' comment('池子备注') VARCHAR(255)"`
}
