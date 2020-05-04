package models

type FranchiseeTeam struct {
	Id        int    `xorm:"not null pk autoincr INT"`
	FranId    int    `xorm:"index INT"`
	UserRole  string `xorm:"VARCHAR(50)"`
	Name      string `xorm:"VARCHAR(50)"`
	Phone     string `xorm:"VARCHAR(50)"`
	UserImg   string `xorm:"VARCHAR(256)"`
	CardFront string `xorm:"VARCHAR(256)"`
	CardBack  string `xorm:"VARCHAR(256)"`
	HoldCard  string `xorm:"VARCHAR(256)"`
	Datils    string `xorm:"VARCHAR(512)"`
}
