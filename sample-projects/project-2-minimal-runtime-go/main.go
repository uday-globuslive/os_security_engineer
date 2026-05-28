package main

import (
    "encoding/json"
    "net/http"
)

type healthResponse struct {
    Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    _, _ = w.Write([]byte(`{"message":"minimal runtime go service"}`))
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/", rootHandler)
    _ = http.ListenAndServe(":8080", nil)
}
