package http

import (
	"net/http"

	"github.com/rs/cors"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/bins", GetTotalBinsHandler)
	mux.HandleFunc("/create/bin", CreateBin)
	mux.HandleFunc("/delete/bin", DeleteBin)

	mux.HandleFunc("/totalItems", GetItemsHandler)
	mux.HandleFunc("/create/item", CreateItem)
	mux.HandleFunc("/delete/item", DeleteItemHandler)

	mux.HandleFunc("/signin", SignIn)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})
	handler := c.Handler(mux)
	return handler
}
