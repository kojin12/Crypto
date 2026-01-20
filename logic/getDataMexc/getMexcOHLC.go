package getdatamexc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetMexcOHLCStrings(symbol, timeframe, limit string) ([][]string, error) {
	u, err := url.Parse("https://api.mexc.com/api/v3/klines")
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("interval", timeframe)
	params.Add("limit", limit)
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var raw [][]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, err
	}

	result := make([][]string, len(raw))
	for i, row := range raw {
		result[i] = make([]string, len(row))
		for j, val := range row {
			result[i][j] = fmt.Sprintf("%v", val)
		}
	}

	return result, nil
}
