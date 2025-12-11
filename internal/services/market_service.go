package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type MarketPrice struct {
	Symbol string  `json:"symbol"`
	Open   float64 `json:"open"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

// Estructuras adaptadas EXACTAMENTE al JSON de la API
type APIResponse struct {
	Data struct {
		QuoteSummary struct {
			Result []struct {
				SummaryDetail struct {
					Open struct {
						Raw float64 `json:"raw"`
					} `json:"open"`
					RegularMarketDayHigh struct {
						Raw float64 `json:"raw"`
					} `json:"regularMarketDayHigh"`
					RegularMarketDayLow struct {
						Raw float64 `json:"raw"`
					} `json:"regularMarketDayLow"`
					RegularMarketPreviousClose struct {
						Raw float64 `json:"raw"`
					} `json:"regularMarketPreviousClose"`
					RegularMarketOpen struct {
						Raw float64 `json:"raw"`
					} `json:"regularMarketOpen"`
					RegularMarketVolume struct {
						Raw int64 `json:"raw"`
					} `json:"regularMarketVolume"`
				} `json:"summaryDetail"`
			} `json:"result"`
		} `json:"quoteSummary"`
	} `json:"data"`
}

func GetPriceFromAPI(symbol string) (*MarketPrice, error) {

	url := fmt.Sprintf("https://%s/v1/stock/summary?symbol=%s",
		os.Getenv("RAPIDAPI_HOST"), symbol)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("RAPIDAPI_HOST"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var data APIResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	sd := data.Data.QuoteSummary.Result[0].SummaryDetail

	price := &MarketPrice{
		Symbol: symbol,
		Open:   sd.Open.Raw,
		High:   sd.RegularMarketDayHigh.Raw,
		Low:    sd.RegularMarketDayLow.Raw,
		Close:  sd.RegularMarketPreviousClose.Raw,
		Volume: sd.RegularMarketVolume.Raw,
	}

	return price, nil
}
