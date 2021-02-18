package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Stock struct {
	ID         uint            `json:"id" gorm:"primary_key"`
	SymbolId   string          `json:"symbolId"`
	SymbolName string          `json:"symbolName"`
	Price      decimal.Decimal `json:"price" sql:"type:decimal(15,2);"`
	RankTime   time.Time       `json:"rankTime" sql:"index"`
}
