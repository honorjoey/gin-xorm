package models

import (
	"time"
)

type CoinInfo struct {
	CoinId                       int       `xorm:"not null pk autoincr comment('币ID') INT"`
	CoinEnglishName              string    `xorm:"VARCHAR(30)"`
	CoinSymbol                   string    `xorm:"VARCHAR(30)"`
	CoinUrl                      string    `xorm:"VARCHAR(255)"`
	CoinChineseName              string    `xorm:"VARCHAR(30)"`
	CoinContractAddress          string    `xorm:"VARCHAR(255)"`
	CoinDecimal                  int       `xorm:"comment('代币真实小数点后位数') INT"`
	CoinDecimalUsed              int       `xorm:"comment('代币显示小数点后位数') INT"`
	CoinQuotesName               string    `xorm:"VARCHAR(255)"`
	CoinLocation                 string    `xorm:"VARCHAR(255)"`
	CoinPriceRmb                 float64   `xorm:"comment('币种人民币价格') DOUBLE(30,4)"`
	CoinPriceUsdt                string    `xorm:"VARCHAR(255)"`
	CoinFluctuation              string    `xorm:"VARCHAR(255)"`
	CoinPriceUpdatetime          time.Time `xorm:"DATETIME"`
	CoinFluctuationUpdatetime    time.Time `xorm:"DATETIME"`
	CoinPriceRequestpath         string    `xorm:"VARCHAR(255)"`
	CoinPriceJsonStructure       string    `xorm:"VARCHAR(255)"`
	CoinPriceRequestmethod       string    `xorm:"VARCHAR(255)"`
	CoinPriceRequestbody         string    `xorm:"VARCHAR(255)"`
	CoinFluctuationJsonStructure string    `xorm:"VARCHAR(255)"`
	CoinFluctuationRequestpath   string    `xorm:"VARCHAR(255)"`
	CoinFluctuationRequestmethod string    `xorm:"VARCHAR(255)"`
	Isdisplay                    int       `xorm:"default 1 comment('是否显示，1-显示，0-不显示') INT"`
	CoinFluctuationRequestbody   string    `xorm:"VARCHAR(255)"`
	CoinSort                     int       `xorm:"comment('排序，越小，越前') INT"`
	IsSync                       int       `xorm:"comment('是否从交易所同步最新价格') INT"`
}
