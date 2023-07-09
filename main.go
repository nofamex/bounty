package main

import (
	"bounty/bounty"
	"bounty/config"
	"bounty/db"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	err := config.ReadConfig()
	if err != nil {
		panic(err.Error())
	}

	go db.MigrateDatabase()

	err = db.ConnectDatabase()
	if err != nil {
		panic(err.Error())
	}
	defer db.DB.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/debug", middleware.Profiler())
	r.Mount("/bounty", bounty.BountyHandler())

	http.ListenAndServe(fmt.Sprintf(":%s", config.ENVConfig.PORT), r)
}
