package main

import (
	"fmt"
	"github.com/aktivgo-gopkg/extractor/ext"
	"github.com/gorilla/mux"
	"github.com/semichkin-gopkg/conv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type (
	GetMemberListRequest struct {
		WorkspaceID string `json:"workspace_id"`
		Pagination  struct {
			Limit  uint64 `json:"limit"`
			Offset uint64 `json:"offset"`
		} `json:",squash"`
	}

	CreateMemberRequest struct {
		WorkspaceID string `json:"workspace_id"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		Phone       string `json:"phone"`
	}
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/workspace/{workspace_id}/members", func(w http.ResponseWriter, r *http.Request) {
		req, err := ext.Extract[GetMemberListRequest](
			r,
			ext.Units{
				{From: "url", Name: "workspace_id", Required: true},
				{From: "query", Name: "limit", Required: false, DefaultValue: 20},
				{From: "query", Name: "offset", DefaultValue: 0},
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
	}).Methods("GET")

	router.HandleFunc("/workspace/{id}/members", func(w http.ResponseWriter, r *http.Request) {
		req, err := ext.Extract[CreateMemberRequest](
			r,
			ext.Units{
				{From: "url", Name: "workspace_id", Required: true},
				{From: "body", Name: "username", Required: true},
				{From: "body", Name: "email", Required: true},
				{From: "body", Name: "phone", Required: false},
				{From: "body", Name: "image", DefaultValue: "member.png"},
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
