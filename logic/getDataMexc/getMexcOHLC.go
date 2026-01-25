package getdatamexc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetMexcOHLC(symbol, timeframe, limit string) ([][]string, float64) {
	u, err := url.Parse("https://api.mexc.com/api/v3/klines")
	if err != nil {
		panic(err)
	}

	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("interval", timeframe)
	params.Add("limit", limit)
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var raw [][]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		panic(err)
	}

	if len(raw) == 0 {
		return nil, 0
	}

	result := make([][]string, len(raw))
	for i, row := range raw {
		result[i] = make([]string, len(row))
		for j, val := range row {
			if j == 0 || j == 6 {
				ts := int64(val.(float64))
				t := time.Unix(ts/1000, 0)
				result[i][j] = t.Format("2006-01-02 15:04:05")
			} else {
				result[i][j] = fmt.Sprintf("%v", val)
			}
		}
	}

	var priceFloat float64
	if len(result) > 0 && len(result[len(result)-1]) > 4 {
		priceFloat, _ = strconv.ParseFloat(result[len(result)-1][4], 64)
	}

	return result, priceFloat
}
