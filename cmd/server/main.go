package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"keisan-aire/internal/config"
	"keisan-aire/internal/logging"
	"keisan-aire/internal/services"
)

func main() {
	logging.Init()
	cfg := config.Load()

	fmt.Println("🚀 Keisan Aire iniciado en modo:", cfg.Env)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})
	http.HandleFunc("/price/", func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Path[len("/price/"):]

		price, err := services.GetPriceFromAPI(symbol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(price)
	})

	fmt.Printf("Servidor escuchando en http://localhost:%s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
