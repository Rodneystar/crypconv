package cc

import (
	"fmt"
	"math"

	"github.com/piquette/finance-go/quote"
)

func Convert(amount float64, coin string) (float64, error) {
	coinDeets := ParseCoin(coin)
	q, err := quote.Get(coinDeets.Ticker + "-USD")

	if err != nil || q == nil {
		return 0, fmt.Errorf("bad coin ticker %s", coin)
	}
	return (amount * q.RegularMarketPrice) * math.Pow(10, float64(coinDeets.Exp)), err
}

type UnitDeets struct {
	Unit   string
	Ticker string
	Exp    int //exponent
}

var unitMap = map[string]UnitDeets{
	"WEI":    {Unit: "WEI", Ticker: "ETH", Exp: -18},
	"KWEI":   {Unit: "KWEI", Ticker: "ETH", Exp: -15},
	"ADA":    {Unit: "ADA", Ticker: "ETH", Exp: -15},
	"MWEI":   {Unit: "MWEI", Ticker: "ETH", Exp: -15},
	"GWEI":   {Unit: "GWEI", Ticker: "ETH", Exp: -9},
	"SZABO":  {Unit: "SZABO", Ticker: "ETH", Exp: -6},
	"FINNEY": {Unit: "FINNEY", Ticker: "ETH", Exp: -3},
	"ETHER":  {Unit: "ETHER", Ticker: "ETH", Exp: 0},
	"KETHER": {Unit: "KETHER", Ticker: "ETH", Exp: 3},
	"METHER": {Unit: "METHER", Ticker: "ETH", Exp: 6},
	"GETHER": {Unit: "GETHER", Ticker: "ETH", Exp: 9},
	"TETHER": {Unit: "TETHER", Ticker: "ETH", Exp: 12},
}

func ParseCoin(coin string) UnitDeets {
	deets, found := unitMap[coin]
	if !found {
		return UnitDeets{Unit: coin, Ticker: coin, Exp: 0}
	}
	return deets
}
