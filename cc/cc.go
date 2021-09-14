package cc

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func Convert(amount float64, coin string) (float64, error) {
	q, err := quote.Get(coin + "-USD")
	if err != nil || q == nil {
		return 0, fmt.Errorf("bad coin ticker %s", coin)
	}
	return amount * q.RegularMarketPrice, err
}
