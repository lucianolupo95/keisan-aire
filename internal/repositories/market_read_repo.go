package repositories

import (
	"context"
	"database/sql"
	"time"
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
func (r *MarketRepository) GetMarketDataStats(assetID int) (count int, lastTimestamp *string, err error) {
	query := `
		SELECT COUNT(*), MAX(timestamp)
		FROM market_data
		WHERE asset_id = $1
	`

	row := r.DB.QueryRow(query, assetID)

	var ts sql.NullTime
	if err := row.Scan(&count, &ts); err != nil {
		return 0, nil, err
	}

	if ts.Valid {
		s := ts.Time.Format(time.RFC3339)
		return count, &s, nil
	}

	return count, nil, nil
}
