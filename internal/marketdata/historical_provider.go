package marketdata

import "time"

type Candle struct {
	Timestamp time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    float64
}

type HistoricalProvider interface {
	FetchDaily(symbol string) ([]Candle, error)
}
