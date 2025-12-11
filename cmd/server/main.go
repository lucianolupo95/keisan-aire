package main

import (
	"fmt"
	"log"
	"net/http"

	"keisan-aire/internal/config"
	"keisan-aire/internal/logging"
)

func main() {
	logging.Init()
	cfg := config.Load()

	fmt.Println("🚀 Keisan Aire iniciado en modo:", cfg.Env)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})

	fmt.Printf("Servidor escuchando en http://localhost:%s\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatal(err)
	}
}
