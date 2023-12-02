package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/matheus-alpe/go-coin-api/internal/handlers"
	"github.com/sirupsen/logrus"
)

func main() {
    logrus.SetReportCaller(true)
    r := chi.NewRouter()
    handlers.Handler(r)

    fmt.Println("Starting GO API")
    err := http.ListenAndServe("localhost:3232", r)
    if err != nil {
        logrus.Error(err)
    }
}
