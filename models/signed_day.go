package models

type SignedDay struct {
	Id     int     `xorm:"not null pk autoincr INT"`
	Days   int     `xorm:"INT"`
	Reward float64 `xorm:"DOUBLE(16,8)"`
}
