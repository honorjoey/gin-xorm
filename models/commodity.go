package models

import (
	"time"
)

type Commodity struct {
	Id             int       `xorm:"not null pk autoincr INT"`
	BrandId        int       `xorm:"INT"`
	Name           string    `xorm:"VARCHAR(50)"`
	Price          float64   `xorm:"DOUBLE(16,4)"`
	GviPrice       float64   `xorm:"DOUBLE(16,4)"`
	Stock          int       `xorm:"INT"`
	Logos          string    `xorm:"VARCHAR(256)"`
	Img            string    `xorm:"VARCHAR(256)"`
	Describes      string    `xorm:"VARCHAR(512)"`
	Details        string    `xorm:"VARCHAR(512)"`
	Explains       string    `xorm:"VARCHAR(512)"`
	Express        string    `xorm:"VARCHAR(512)"`
	Freight        float64   `xorm:"DOUBLE(16,4)"`
	Status         int       `xorm:"INT"`
	Creator        string    `xorm:"VARCHAR(256)"`
	CreateTime     time.Time `xorm:"DATETIME"`
	LastUpdateTime time.Time `xorm:"DATETIME"`
	Units          string    `xorm:"VARCHAR(50)"`
	TimeValidity   string    `xorm:"VARCHAR(512)"`
	Discount       string    `xorm:"VARCHAR(255)"`
}
