package main

import (
	"fmt"
	"github.com/aktivgo-gopkg/extractor"
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
		ID    string `json:"id" httpext:"from:url,name:id,required"`
		Title string `json:"title" httpext:"from:body,name:title,required"`
		Logo  string `json:"logo" httpext:"from:body,name:logo,default:logo"`
	}
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/example/{id}", func(w http.ResponseWriter, r *http.Request) {
		req, err := extractor.Extract[Request](r)
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
