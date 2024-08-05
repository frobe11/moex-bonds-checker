package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/frobe11/moex-bonds-checker/iternal/models"
)

const (
	baseUrl = "https://iss.moex.com/iss/"
)

func Getbonds() ([]Bond, error) {
	url := fmt.Sprintf("%s%s", baseUrl,
		"engines/stock/markets/bonds/securities.json?iss.only=securities,marketdata&lang=en&iss.meta=off&iss.json=extended")
	fmt.Print(url)
	res := make([]Bond, 0)
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}

	b := new(bondInfo)
	fmt.Print(*resp)
	err = getJson(*resp, &b)
	if err != nil {
		return res, err
	}
	fmt.Print(b)
	parseBonds(*b, res);
	fmt.Print(res);	
	return res, nil
}

type bondInfo struct {
	Charsetinfo struct {
		Name string `json:"name"`
	} `json:"charsetinfo,omitempty"`

	Securities []struct {
		Secid                 string      `json:"SECID"`
		Boardid               string      `json:"BOARDID"`
		Shortname             string      `json:"SHORTNAME"`
		Prevwaprice           float64     `json:"PREVWAPRICE"`
		Yieldatprevwaprice    float64     `json:"YIELDATPREVWAPRICE"`
		Couponvalue           float64     `json:"COUPONVALUE"`
		Nextcoupon            string      `json:"NEXTCOUPON"`
		Accruedint            float64     `json:"ACCRUEDINT"`
		Prevprice             float64     `json:"PREVPRICE"`
		Lotsize               int         `json:"LOTSIZE"`
		Facevalue             float64     `json:"FACEVALUE"`
		Boardname             string      `json:"BOARDNAME"`
		Status                string      `json:"STATUS"`
		Matdate               string      `json:"MATDATE"`
		Decimals              int         `json:"DECIMALS"`
		Couponperiod          int         `json:"COUPONPERIOD"`
		Issuesize             int         `json:"ISSUESIZE"`
		Prevlegalcloseprice   float64     `json:"PREVLEGALCLOSEPRICE"`
		Prevdate              string      `json:"PREVDATE"`
		Secname               string      `json:"SECNAME"`
		Remarks               interface{} `json:"REMARKS"`
		Marketcode            string      `json:"MARKETCODE"`
		Instrid               string      `json:"INSTRID"`
		Sectorid              interface{} `json:"SECTORID"`
		Minstep               float64     `json:"MINSTEP"`
		Faceunit              string      `json:"FACEUNIT"`
		Buybackprice          float64     `json:"BUYBACKPRICE"`
		Buybackdate           string      `json:"BUYBACKDATE"`
		Isin                  string      `json:"ISIN"`
		Latname               string      `json:"LATNAME"`
		Regnumber             string      `json:"REGNUMBER"`
		Currencyid            string      `json:"CURRENCYID"`
		Issuesizeplaced       int         `json:"ISSUESIZEPLACED"`
		Listlevel             int         `json:"LISTLEVEL"`
		Sectype               int         `json:"SECTYPE"`
		Couponpercent         float64     `json:"COUPONPERCENT"`
		Offerdate             string      `json:"OFFERDATE"`
		Settledate            string      `json:"SETTLEDATE"`
		Lotvalue              float64     `json:"LOTVALUE"`
		Facevalueonsettledate float64     `json:"FACEVALUEONSETTLEDATE"`
	} `json:"securities,omitempty"`

	Marketdata []struct {
		Secid                 string      `json:"SECID"`
		Bid                   interface{} `json:"BID"`
		Biddepth              interface{} `json:"BIDDEPTH"`
		Offer                 interface{} `json:"OFFER"`
		Offerdepth            interface{} `json:"OFFERDEPTH"`
		Spread                float64     `json:"SPREAD"`
		Biddeptht             interface{} `json:"BIDDEPTHT"`
		Offerdeptht           interface{} `json:"OFFERDEPTHT"`
		Open                  float64     `json:"OPEN"`
		Low                   float64     `json:"LOW"`
		High                  float64     `json:"HIGH"`
		Last                  float64     `json:"LAST"`
		Lastchange            float64     `json:"LASTCHANGE"`
		Lastchangeprcnt       float64     `json:"LASTCHANGEPRCNT"`
		Qty                   int         `json:"QTY"`
		Value                 float64     `json:"VALUE"`
		Yield                 float64     `json:"YIELD"`
		ValueUsd              float64     `json:"VALUE_USD"`
		Waprice               float64     `json:"WAPRICE"`
		Lastcngtolastwaprice  float64     `json:"LASTCNGTOLASTWAPRICE"`
		Waptoprevwapriceprcnt float64     `json:"WAPTOPREVWAPRICEPRCNT"`
		Waptoprevwaprice      float64     `json:"WAPTOPREVWAPRICE"`
		Yieldatwaprice        float64     `json:"YIELDATWAPRICE"`
		Yieldtoprevyield      float64     `json:"YIELDTOPREVYIELD"`
		Closeyield            float64     `json:"CLOSEYIELD"`
		Closeprice            interface{} `json:"CLOSEPRICE"`
		Marketpricetoday      float64     `json:"MARKETPRICETODAY"`
		Marketprice           float64     `json:"MARKETPRICE"`
		Lasttoprevprice       float64     `json:"LASTTOPREVPRICE"`
		Numtrades             int         `json:"NUMTRADES"`
		Voltoday              float64     `json:"VOLTODAY"`
		Valtoday              float64     `json:"VALTODAY"`
		ValtodayUsd           float64     `json:"VALTODAY_USD"`
		Boardid               string      `json:"BOARDID"`
		Tradingstatus         string      `json:"TRADINGSTATUS"`
		Updatetime            string      `json:"UPDATETIME"`
		Dturation             float64     `json:"DURATION"`
		Numbids               interface{} `json:"NUMBIDS"`
		Numoffers             interface{} `json:"NUMOFFERS"`
		Change                float64     `json:"CHANGE"`
		Time                  string      `json:"TIME"`
		Highbid               interface{} `json:"HIGHBID"`
		Lowoffer              interface{} `json:"LOWOFFER"`
		Priceminusprevwaprice float64     `json:"PRICEMINUSPREVWAPRICE"`
		LastBid               interface{} `json:"LASTBID"`
		Lastoffer             interface{} `json:"LASTOFFER"`
		Lcurrentprice         float64     `json:"LCURRENTPRICE"`
		Lcloseprice           float64     `json:"LCLOSEPRICE"`
		Marketprice2          interface{} `json:"MARKETPRICE2"`
		Openperiodprice       interface{} `json:"OPENPERIODPRICE"`
		Seqnum                int64       `json:"SEQNUM"`
		Systime               string      `json:"SYSTIME"`
		ValtodayRur           int64       `json:"VALTODAY_RUR"`
		Iricpiclose           interface{} `json:"IRICPICLOSE"`
		Beiclose              interface{} `json:"BEICLOSE"`
		Cbrclose              interface{} `json:"CBRCLOSE"`
		Yieldtooffer          interface{} `json:"YIELDTOOFFER"`
		Yieldlastcoupon       interface{} `json:"YIELDLASTCOUPON"`
		Tradingsession        interface{} `json:"TRADINGSESSION"`
	} `json:"marketdata,omitempty"`
}

func getJson(r http.Response, target interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func parseBonds(bInfo bondInfo, bTarget []Bond) {
	for i, s := range bInfo.Securities {
		bTarget = append(bTarget,
			*New(s.Shortname,
				s.Couponpercent,
				s.Couponvalue,
				bInfo.Marketdata[i].Last,
				s.Facevalue,
				s.Accruedint,
				s.Nextcoupon))
	}
}
