package routes

import (
	"main/handlers"
	"net/http"
)

func GetNewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/data/info", handlers.GetDataHandlers)
	mux.HandleFunc("/data/price_change", handlers.GetPriceChangeHandlers)
	return mux
}
