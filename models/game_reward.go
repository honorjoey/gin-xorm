package models

import (
	"time"
)

type GameReward struct {
	Id          int       `xorm:"not null pk autoincr INT"`
	MemberId    int       `xorm:"comment('用户id') INT"`
	UserAccount string    `xorm:"VARCHAR(255)"`
	UserAsset   float64   `xorm:"comment('用户消耗资产') DOUBLE(18,8)"`
	UserScore   int       `xorm:"comment('用户游戏得分') INT"`
	GameId      int       `xorm:"comment('游戏id') INT"`
	RankLevel   int       `xorm:"comment('用户奖励等级') INT"`
	UserIndex   int       `xorm:"comment('用户该游戏等级') INT"`
	Rank        float64   `xorm:"comment('用户该游戏的奖励') DOUBLE(18,8)"`
	RankPool    float64   `xorm:"comment('该游戏的奖金池') DOUBLE(18,8)"`
	CreateTime  time.Time `xorm:"comment('创建时间') DATETIME"`
}
