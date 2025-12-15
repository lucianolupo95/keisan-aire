package services

import (
	"time"

	"keisan-aire/internal/repositories"
)

type HistoricalPrice struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

// ⚠️ Por ahora mock: en la próxima instrucción conectamos API real
func FetchHistoricalPrices(symbol string, days int) ([]HistoricalPrice, error) {
	var prices []HistoricalPrice

	now := time.Now()
	for i := days; i > 0; i-- {
		prices = append(prices, HistoricalPrice{
			Timestamp: now.AddDate(0, 0, -i),
			Open:      100 + float64(i),
			High:      101 + float64(i),
			Low:       99 + float64(i),
			Close:     100 + float64(i)/2,
			Volume:    1000000,
		})
	}

	return prices, nil
}

func LoadHistoricalIfNeeded(
	repo *repositories.MarketRepository,
	assetID int,
	symbol string,
) error {

	count, _, err := repo.GetMarketDataStats(assetID)
	if err != nil {
		return err
	}

	if count >= 600 {
		return nil
	}

	history, err := FetchHistoricalPrices(symbol, 700)
	if err != nil {
		return err
	}

	for _, h := range history {
		err := repo.InsertMarketDataWithTimestamp(
			assetID,
			h.Open,
			h.High,
			h.Low,
			h.Close,
			h.Volume,
			h.Timestamp,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
