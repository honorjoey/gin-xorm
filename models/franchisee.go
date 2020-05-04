package models

import (
	"time"
)

type Franchisee struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	FranName       string    `xorm:"VARCHAR(50)"`
	FranNum        string    `xorm:"VARCHAR(50)"`
	BrandId        int       `xorm:"INT"`
	Period         int       `xorm:"INT"`
	Address        string    `xorm:"VARCHAR(256)"`
	Quota          int       `xorm:"INT"`
	ExTurnover     string    `xorm:"VARCHAR(256)"`
	ReCycle        int       `xorm:"INT"`
	ExIncome       string    `xorm:"VARCHAR(256)"`
	Details        string    `xorm:"VARCHAR(512)"`
	ImgBanner      string    `xorm:"VARCHAR(512)"`
	ReleaseTime    time.Time `xorm:"DATETIME"`
	EndTime        time.Time `xorm:"DATETIME"`
	Status         int       `xorm:"INT"`
	ReleaseStatus  int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LsatUpdateTime time.Time `xorm:"DATETIME"`
	UserId         int       `xorm:"INT"`
}
