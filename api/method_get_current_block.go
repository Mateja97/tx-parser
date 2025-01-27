package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tx-parser/parser"
)

func GetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	currentBlock := parser.GetCurrentBlock()

	response := map[string]string{"message": fmt.Sprintf("current block: %d", currentBlock)}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Printf("error encoding subscribe response: %v", err)
	}
}
