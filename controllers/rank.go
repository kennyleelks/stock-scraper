package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"log"
	"main/models"
	"net/http"
	"time"
)

type resp struct {
	Data data `json:"data"`
}

type data struct {
	Stocks     []stock    `json:"list"`
	Pagination pagination `json:"pagination"`
	RankTime   string     `json:"rankTime"`
}

type pagination struct {
	ResultsTotal int `json:"resultsTotal"`
	NextOffset   int `json:"nextOffset,string"`
}

type stock struct {
	Symbol     string          `json:"symbol"`
	SymbolName string          `json:"symbolName"`
	Price      decimal.Decimal `json:"price,string"`
}

func GetRankList(offset int) {
	url := fmt.Sprintf("https://tw.stock.yahoo.com/_td/api/resource/StockServices.rank;exchange=ALL;limit=100;offset=%d;period=1D;sortBy=-volume?bkt=tw-qsp-exp&device=desktop&ecma=default&feature=&intl=tw&lang=zh-Hant-TW&partner=none&prid=577nuahg2pm8h&region=TW&site=finance&tz=Asia%%2FTaipei&ver=1.2.734&returnMeta=true", offset)
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var resp resp
	if err := json.Unmarshal(body, &resp); err != nil {
		panic(err.Error())
	}

	for _, t := range resp.Data.Stocks {
		fmt.Println(t)

		rankTime, err := time.Parse(time.RFC3339, resp.Data.RankTime)
		if err != nil {
			log.Println(err.Error())
		} else {
			stock := models.Stock{
				SymbolId:   t.Symbol,
				SymbolName: t.SymbolName,
				Price:      t.Price,
				RankTime:   rankTime,
			}

			if models.DB.Model(&models.Stock{}).Where("symbol_id = ?", t.Symbol).Updates(models.Stock{Price: t.Price, RankTime: rankTime}).RowsAffected == 0 {
				models.DB.Create(&stock)
			}
		}
	}

	nextOffset := resp.Data.Pagination.NextOffset
	if nextOffset > 0 {
		GetRankList(nextOffset)
	}
}
