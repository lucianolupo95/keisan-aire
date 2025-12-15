package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"keisan-aire/internal/analysis"
	"keisan-aire/internal/config"
	"keisan-aire/internal/domain"
	"keisan-aire/internal/indicators"
	"keisan-aire/internal/logging"
	"keisan-aire/internal/repositories"
	"keisan-aire/internal/services"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// --- Init ---
	logging.Init()
	cfg := config.Load()

	fmt.Println("🚀 Keisan Aire iniciado en modo:", cfg.Env)

	// --- DB ---
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error al abrir DB:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("No se pudo conectar a PostgreSQL:", err)
	}
	fmt.Println("📦 Conectado a PostgreSQL")

	repo := repositories.NewMarketRepository(db)

	// --- Health ---
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})

	// --- Price ---
	http.HandleFunc("/price/", func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Path[len("/price/"):]

		price, err := services.GetPriceFromAPI(symbol)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		assetID := 1 // TODO: resolver con tabla assets

		if err := repo.InsertMarketData(
			assetID,
			price.Open,
			price.High,
			price.Low,
			price.Close,
			price.Volume,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(price)
	})

	// --- Analysis ---
	http.HandleFunc("/analysis/", func(w http.ResponseWriter, r *http.Request) {

		symbol := r.URL.Path[len("/analysis/"):]
		assetID := 1 // temporal
		// --- Carga histórica automática ---
		if err := services.LoadHistoricalIfNeeded(repo, assetID, symbol); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// --- Datos reales desde DB ---
		pricesShort, err := repo.GetLastCloses(assetID, 100)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pricesMedium, err := repo.GetLastCloses(assetID, 300)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pricesLong, err := repo.GetLastCloses(assetID, 600)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// --- Indicadores ---
		shortFastSMA := indicators.SMA(pricesShort, 9)
		shortSlowSMA := indicators.SMA(pricesShort, 21)

		mediumFastSMA := indicators.SMA(pricesMedium, 20)
		mediumSlowSMA := indicators.SMA(pricesMedium, 50)

		longFastSMA := indicators.SMA(pricesLong, 50)
		longSlowSMA := indicators.SMA(pricesLong, 200)

		shortStochastic := indicators.StochasticK(pricesShort, 14)

		// --- Señales ---
		shortSignals := []domain.IndicatorSignal{
			analysis.SMACrossSignal(shortFastSMA, shortSlowSMA, domain.TimeframeShort),
			analysis.StochasticSignal(shortStochastic),
		}

		mediumSignals := []domain.IndicatorSignal{
			analysis.SMACrossSignal(mediumFastSMA, mediumSlowSMA, domain.TimeframeMedium),
		}

		longSignals := []domain.IndicatorSignal{
			analysis.SMACrossSignal(longFastSMA, longSlowSMA, domain.TimeframeLong),
		}

		// --- Agregación ---
		result := domain.AssetAnalysisResponse{
			Symbol: symbol,
			Short:  analysis.AggregateSignals(domain.TimeframeShort, shortSignals),
			Medium: analysis.AggregateSignals(domain.TimeframeMedium, mediumSignals),
			Long:   analysis.AggregateSignals(domain.TimeframeLong, longSignals),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	handler := withCORS(http.DefaultServeMux)

	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatal(err)
	}

}
