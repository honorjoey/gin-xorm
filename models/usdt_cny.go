package models

type UsdtCny struct {
	Id        int     `xorm:"not null pk autoincr INT"`
	Usdt      float64 `xorm:"DOUBLE(10,4)"`
	Timestamp int     `xorm:"INT"`
}
