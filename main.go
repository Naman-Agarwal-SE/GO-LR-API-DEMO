package main

import (
	"golang-demo/helper"
	"golang-demo/router"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("started server!")
	routes := router.NewRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
