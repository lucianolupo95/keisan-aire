package repositories

import (
	"context"
	"database/sql"
	"time"
)

type MarketRepository struct {
	DB *sql.DB
}

func NewMarketRepository(db *sql.DB) *MarketRepository {
	return &MarketRepository{DB: db}
}

func (r *MarketRepository) InsertMarketData(assetID int, open, high, low, close float64, volume int64) error {
	query := `
        INSERT INTO market_data (asset_id, timestamp, open, high, low, close, volume)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `
	_, err := r.DB.ExecContext(
		context.Background(),
		query,
		assetID,
		time.Now().UTC(),
		open, high, low, close, volume,
	)
	return err
}
