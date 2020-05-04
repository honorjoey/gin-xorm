package models

type Constant struct {
	Id           int    `xorm:"not null pk autoincr INT"`
	ItemKey      string `xorm:"VARCHAR(255)"`
	ItemValue    string `xorm:"VARCHAR(255)"`
	ItemDescribe string `xorm:"VARCHAR(255)"`
	Name         string `xorm:"VARCHAR(255)"`
}
