package main

import (
    "encoding/json"
    "net/http"
    "time"
)

type TimeResponse struct {
    Time string `json:"time"`
}

func main() {
    http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
        currentTime := time.Now().Format(time.RFC3339)
        response := TimeResponse{Time: currentTime}
        jsonResponse, err := json.Marshal(response)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write(jsonResponse)
    })

    // Запуск веб-сервера на порту 8795
    if err := http.ListenAndServe(":8795", nil); err != nil {
        panic(err)
    }
}