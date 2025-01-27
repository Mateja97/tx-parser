package api

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"log"
	"net/http"
	"tx-parser/parser"
)

type SubscribeRequest struct {
	Address string `json:"address"`
}

func Subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req SubscribeRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("closing body error: %v", err)
		}
	}(r.Body)

	if !common.IsHexAddress(req.Address) {
		http.Error(w, "Invalid ethereum address", http.StatusBadRequest)
		return
	}
	ok := parser.Subscribe(req.Address)
	if !ok {
		response := map[string]string{"message": "Failed to subscribe"}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}

	response := map[string]string{"message": "Successfully subscribed"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error encoding subscribe response: %v", err)
	}
}
