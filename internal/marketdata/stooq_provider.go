package marketdata

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type StooqProvider struct {
	HTTPClient *http.Client
}

func NewStooqProvider(client *http.Client) *StooqProvider {
	if client == nil {
		client = &http.Client{Timeout: 15 * time.Second}
	}
	return &StooqProvider{HTTPClient: client}
}

// Stooq uses lowercase symbols; for US stocks commonly: aapl.us
func (p *StooqProvider) FetchDaily(symbol string) ([]Candle, error) {
	s := normalizeStooqSymbol(symbol)
	url := fmt.Sprintf("https://stooq.com/q/d/l/?s=%s&i=d", s)

	resp, err := p.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("stooq status %d", resp.StatusCode)
	}

	r := csv.NewReader(resp.Body)
	r.FieldsPerRecord = -1

	// header: Date,Open,High,Low,Close,Volume
	_, err = r.Read()
	if err != nil {
		return nil, fmt.Errorf("csv header: %w", err)
	}

	var out []Candle
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil || len(rec) < 6 {
			continue
		}

		ts, err := time.Parse("2006-01-02", rec[0])
		if err != nil {
			continue
		}

		open, _ := strconv.ParseFloat(rec[1], 64)
		high, _ := strconv.ParseFloat(rec[2], 64)
		low, _ := strconv.ParseFloat(rec[3], 64)
		closep, _ := strconv.ParseFloat(rec[4], 64)
		vol, _ := strconv.ParseFloat(rec[5], 64)

		out = append(out, Candle{
			Timestamp: ts,
			Open:      open,
			High:      high,
			Low:       low,
			Close:     closep,
			Volume:    vol,
		})
	}

	if len(out) == 0 {
		return nil, fmt.Errorf("no candles returned for %s", s)
	}

	return out, nil
}

func normalizeStooqSymbol(symbol string) string {
	s := strings.TrimSpace(strings.ToLower(symbol))
	// si te pasan "AAPL" => "aapl.us" (mínimo viable)
	if !strings.Contains(s, ".") {
		s = s + ".us"
	}
	return s
}
