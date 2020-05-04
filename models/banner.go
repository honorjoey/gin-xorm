package models

import (
	"time"
)

type Banner struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	ImgName        string    `xorm:"VARCHAR(256)"`
	Img            string    `xorm:"VARCHAR(256)"`
	Url            string    `xorm:"VARCHAR(256)"`
	ShowPosition   string    `xorm:"VARCHAR(256)"`
	Creator        string    `xorm:"VARCHAR(256)"`
	Status         int       `xorm:"INT"`
	BannerNum      int       `xorm:"INT"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
}
