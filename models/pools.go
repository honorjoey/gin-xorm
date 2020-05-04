package models

type Pools struct {
	Id          int     `xorm:"not null pk autoincr INT"`
	CoinId      int     `xorm:"not null comment('币ID') INT"`
	PoolId      int     `xorm:"not null comment('池子ID') INT"`
	PoolAddress string  `xorm:"not null comment('池子地址') VARCHAR(255)"`
	PoolPrivkey string  `xorm:"comment('池子私钥') VARCHAR(255)"`
	Scale       float64 `xorm:"not null default 0.0000 comment('池子份额') DOUBLE(16,4)"`
	PoolRemarks string  `xorm:"default '' comment('池子备注') VARCHAR(255)"`
}
