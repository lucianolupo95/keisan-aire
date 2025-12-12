package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"keisan-aire/internal/config"
	"keisan-aire/internal/logging"
	"keisan-aire/internal/repositories"
	"keisan-aire/internal/services"
)

func main() {
	logging.Init()
	cfg := config.Load()

	fmt.Println("🚀 Keisan Aire iniciado en modo:", cfg.Env)

	// Conexión a PostgreSQL
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error al abrir DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("No se pudo conectar a PostgreSQL:", err)
	}
	fmt.Println("📦 Conectado a PostgreSQL")

	repo := repositories.NewMarketRepository(db)

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

		assetID := 1 // TODO reemplazar por tabla assets

		if err := repo.InsertMarketData(assetID, price.Open, price.High, price.Low, price.Close, price.Volume); err != nil {
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
