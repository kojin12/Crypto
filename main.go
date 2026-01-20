package main

import (
	"fmt"
	alert "main/alert"
	logic "main/logic"
	getanalispair "main/logic/getAnalisPair"
	getdatabybit "main/logic/getDataBybit"
)

func main() {
	data, price := getdatabybit.GetOHLC("15", "TONUSDT", "200")
	data60, _ := getdatabybit.GetOHLC("60", "TONUSDT", "200")
	res := logic.Final(data, data60)
	sc := logic.GetScore(res, price)
	pair := alert.GetCoinList()
	btc := getanalispair.GetPriceChange("BTCUSDT")
	fmt.Println(res, sc, price, pair, btc)

}
