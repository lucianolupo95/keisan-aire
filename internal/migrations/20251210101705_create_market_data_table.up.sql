CREATE TABLE IF NOT EXISTS market_data (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL,
    open NUMERIC(18,6),
    high NUMERIC(18,6),
    low NUMERIC(18,6),
    close NUMERIC(18,6),
    volume NUMERIC(30,6),

    -- Relaciones
    CONSTRAINT fk_asset
        FOREIGN KEY (asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

-- Índices recomendados para rendimiento
CREATE INDEX IF NOT EXISTS idx_market_data_asset_timestamp
ON market_data (asset_id, timestamp);
