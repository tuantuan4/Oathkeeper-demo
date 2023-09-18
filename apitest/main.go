package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ok", okHandler)
	http.HandleFunc("/forbidden", forbiddenHandler)
	http.ListenAndServe(":9999", nil)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"message": "Invalid JSON"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Request processed successfully", "data": %v}`, data)
}

func forbiddenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if !isAuthenticated(r) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, `{"message": "Access Forbidden"}`)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Access Granted"}`)
}

func isAuthenticated(req *http.Request) bool {
	return req.Header.Get("Authorization") != ""
}
