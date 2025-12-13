package repositories

import (
	"context"
)

func (r *MarketRepository) GetLastCloses(assetID int, limit int) ([]float64, error) {
	query := `
		SELECT close
		FROM market_data
		WHERE asset_id = $1
		ORDER BY timestamp ASC
		LIMIT $2
	`

	rows, err := r.DB.QueryContext(context.Background(), query, assetID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var closes []float64
	for rows.Next() {
		var c float64
		if err := rows.Scan(&c); err != nil {
			return nil, err
		}
		closes = append(closes, c)
	}

	return closes, nil
}
