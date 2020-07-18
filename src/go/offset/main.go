package main

import (
 	"fmt"
	"github.com/luno/luno-go/decimal"
)

func main() {

	var pastAsks []decimal.Decimal

	// initialising values within bot portfolio
	tradingPeriod := int64(14)
//	StopLossMultDecimal := decimal.NewFromFloat64(0.999, 8)
	offset, _ := decimal.NewFromString("0.00000015")

	// initialising bot
	offsetBot := offsetBot{
		tradingPeriod:	14,								// How often the bot calculates a long term result
	  ema: 				     	decimal.Zero(),         // exponentially smoothed Wilder's MMA for upward change
	  currRow:          16,                  // current row within excel spreadsheet
		offset:						offset,
		readyToBuy:				true,
	}

	initialiseFunds(decimal.NewFromFloat64(0.014,8), decimal.Zero())

	var i int64
	for i = 0; i < tradingPeriod; i++ {
		pastAsks = append(pastAsks, getOfflineAsk(i+1))

	}

	fmt.Println("\n\n\n OFFSET:",offsetBot.offset, "\n")
	offsetBot.ema = sma(pastAsks)
	for i:= 0; i < 1300; i ++ {
		offsetBot.trade()
	}

}
