CREATE TABLE IF NOT EXISTS fundamental_data (
    id SERIAL PRIMARY KEY,
    asset_id INTEGER NOT NULL,

    -- Datos fundamentales
    pe_ratio NUMERIC(18,6),
    eps NUMERIC(18,6),
    market_cap NUMERIC(30,2),
    dividend_yield NUMERIC(18,6),
    debt_to_equity NUMERIC(18,6),
    revenue_growth NUMERIC(18,6),
    profit_margin NUMERIC(18,6),
    sector VARCHAR(100),
    industry VARCHAR(100),

    timestamp TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_asset_fundamental
        FOREIGN KEY (asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_fundamental_asset_timestamp
ON fundamental_data (asset_id, timestamp);
