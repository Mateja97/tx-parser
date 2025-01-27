package api

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"io"
	"log"
	"net/http"
	"tx-parser/parser"
)

type GetTransactionsRequest struct {
	Address string `json:"address"`
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GetTransactionsRequest
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
	
	transactions := parser.GetTransactions(req.Address)
	if len(transactions) == 0 {
		response := map[string]string{"message": "transactions not found"}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("error encoding response: %v", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]any{"message": transactions}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error encoding get transactions response: %v", err)
	}
}
