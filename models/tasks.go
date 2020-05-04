package models

type Tasks struct {
	Id     int    `xorm:"not null pk autoincr INT"`
	Name   string `xorm:"VARCHAR(256)"`
	Datils string `xorm:"VARCHAR(512)"`
	Status int    `xorm:"INT"`
}
