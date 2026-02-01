package main

import (
	"fmt"
	logic "main/logic"
	getdatamexc "main/logic/getDataMexc"
)

func main() {
	//ZORA
	zoraOHLC, price := getdatamexc.GetMexcOHLC("BULLAUSDT", "15m", "200")
	zoraOHLC60, _ := getdatamexc.GetMexcOHLC("BULLAUSDT", "60m", "200")

	res := logic.Final(zoraOHLC, zoraOHLC60)
	sc := logic.GetScore(res, price)

	fmt.Println("Result", price, res, sc)

}
