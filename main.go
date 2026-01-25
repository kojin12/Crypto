package main

import (
	"fmt"

	alert "main/alert"
	logic "main/logic"
	getdatamexc "main/logic/getDataMexc"
)

func main() {
	// Получаем OHLC данные
	ohlc, price := getdatamexc.GetMexcOHLC("SOLUSDT", "1m", "200")
	ohlc60, _ := getdatamexc.GetMexcOHLC("SOLUSDT", "60m", "200")

	// Анализируем свечи
	data := logic.Final(ohlc, ohlc60)
	sc := logic.GetScore(data, price)

	fmt.Println("Candle Analysis:", data)
	fmt.Println("Current Price:", price)
	fmt.Println("Score:", sc)

	spotPrice := alert.GetSpotPrice("SOLUSDT")
	linearPrice := alert.GetLinearPrice("SOL_USDT")

	fmt.Println("Spot Price:", spotPrice)
	fmt.Println("Linear Price:", linearPrice)

	deposit := 1000.0
	change := alert.CalcChangeAndProfit(spotPrice, linearPrice, deposit)

	if change.Result {
		fmt.Printf("Выгодная сделка!\nПрибыль: %.2f\nРазница: %.5f\n", change.Profit, change.Diff)
	} else {
		fmt.Println("Сейчас сделка не выгодна")
	}
}
