package main

import (
	"fmt"
	logic "main/logic"
	getdatamexc "main/logic/getDataMexc"
	"time"
)

func main() {
	zoraOHLC, price := getdatamexc.GetMexcOHLC("BULLAUSDT", "60m", "200")
	time.Sleep(2 * time.Second)
	zoraOHLC60, _ := getdatamexc.GetMexcOHLC("BULLAUSDT", "4h", "200")
	time.Sleep(2 * time.Second)
	res := logic.Final(zoraOHLC, zoraOHLC60)
	sc := logic.GetScore(res, price)

	fmt.Println("Result", price, res, sc)

}
