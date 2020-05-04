package models

type GameAwardLevel struct {
	Id          int     `xorm:"not null pk autoincr INT"`
	GameId      int     `xorm:"comment('游戏id') INT"`
	Level       int     `xorm:"comment('奖金名次') INT"`
	MinIndex    int     `xorm:"comment('最小奖金名次id') INT"`
	MaxIndex    int     `xorm:"comment('最大奖金名次id') INT"`
	RewardCnt   int     `xorm:"comment('该奖分配的人数') INT"`
	RewardRatio float64 `xorm:"comment('奖金分配的比率') DOUBLE(18,6)"`
}
