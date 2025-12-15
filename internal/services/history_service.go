package services

import (
	"keisan-aire/internal/marketdata"
	"keisan-aire/internal/repositories"
	"time"
)

type HistoricalPrice struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

func FetchHistoricalPrices(symbol string) ([]HistoricalPrice, error) {
	provider := marketdata.NewStooqProvider(nil)

	candles, err := provider.FetchDaily(symbol)
	if err != nil {
		return nil, err
	}

	prices := make([]HistoricalPrice, 0, len(candles))
	for _, c := range candles {
		prices = append(prices, HistoricalPrice{
			Timestamp: c.Timestamp,
			Close:     c.Close,
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

	history, err := FetchHistoricalPrices(symbol)
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
