package main

import (
	"fmt"
	alert "main/alert"
	logic "main/logic"
	getdatabybit "main/logic/getDataBybit"
)

func main() {
	data, price := getdatabybit.GetOHLC("15", "TONUSDT", "200")
	data60, _ := getdatabybit.GetOHLC("60", "TONUSDT", "200")

	res := logic.Final(data, data60)
	sc := logic.GetScore(res, price)
	pair := alert.GetCoinList()

	fmt.Println(res, sc, price, pair)
}
