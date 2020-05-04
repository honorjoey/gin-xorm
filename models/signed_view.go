package models

type SignedView struct {
	Id     int `xorm:"not null pk autoincr INT"`
	UserId int `xorm:"INT"`
	Days   int `xorm:"INT"`
}
