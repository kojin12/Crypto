package routes

import (
	"main/handlers"
	"net/http"
)

func GetNewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/data", handlers.GetDataHandlers)
	return mux
}
