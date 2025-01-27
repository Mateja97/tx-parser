package api

import (
	"log"
	"net/http"
)

type handler struct {
	port string
}

var h *handler

func Init(options ...func(*handler)) {
	h = new(handler)
	for _, o := range options {
		o(h)
	}
	http.HandleFunc("/subscribe", Subscribe)
	http.HandleFunc("/getCurrentBlock", GetCurrentBlock)
	http.HandleFunc("/getTransactions", GetTransactions)

}

func ServerRun() error {
	log.Printf("server started on: %s", h.port)
	return http.ListenAndServe(h.port, nil)
}
