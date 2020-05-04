package models

type GameList struct {
	Id         int     `xorm:"not null pk autoincr INT"`
	Name       string  `xorm:"comment('游戏简称') VARCHAR(255)"`
	Amount     float64 `xorm:"comment('游戏消耗的平台币') DOUBLE(18,6)"`
	Ratio      float64 `xorm:"comment('该游戏奖金池的抽成比例') DOUBLE(18,6)"`
	EnableFlag int     `xorm:"comment('该游戏是否使能的标记') INT"`
	PicUrl     string  `xorm:"VARCHAR(255)"`
	GameUrl    string  `xorm:"VARCHAR(255)"`
}
