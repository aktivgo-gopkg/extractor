package main

import (
	"fmt"
	"github.com/aktivgo-gopkg/extractor/http/def"
	"github.com/aktivgo-gopkg/extractor/http/ext"
	"github.com/gorilla/mux"
	"github.com/semichkin-gopkg/conv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	LogoDefaultValue = "logo.png"
)

type (
	Request struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Logo  string `json:"logo"`
	}
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/example/{id}", func(w http.ResponseWriter, r *http.Request) {
		req, err := ext.ExtractUnits[Request](
			r,
			ext.UnitsCollection{
				UrlVars: def.UnitCollection{
					{Name: "id", Required: true},
				},
				Body: def.UnitCollection{
					{Name: "title", Required: true},
					{Name: "logo", DefaultValue: LogoDefaultValue, Required: false},
				},
			},
		)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		js, err := conv.JSON(req)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		_, _ = w.Write(js)
	}).Methods("POST")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}

		defer server.Close()
	}()

	log.Println("server started")

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	_ = <-termChan
}
