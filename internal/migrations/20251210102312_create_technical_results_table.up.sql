CREATE TABLE IF NOT EXISTS technical_results (
    id SERIAL PRIMARY KEY,

    -- Relación con market_data
    market_data_id INTEGER NOT NULL,
    asset_id INTEGER NOT NULL,

    -- Indicadores comunes
    rsi NUMERIC(18,6),
    macd NUMERIC(18,6),
    macd_signal NUMERIC(18,6),
    macd_hist NUMERIC(18,6),
    ma_20 NUMERIC(18,6),
    ma_50 NUMERIC(18,6),
    ma_200 NUMERIC(18,6),
    bollinger_upper NUMERIC(18,6),
    bollinger_middle NUMERIC(18,6),
    bollinger_lower NUMERIC(18,6),
    atr NUMERIC(18,6),

    timestamp TIMESTAMPTZ NOT NULL,

    -- Relaciones
    CONSTRAINT fk_market_data
        FOREIGN KEY (market_data_id)
        REFERENCES market_data(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_asset_technical
        FOREIGN KEY (asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

-- Índice para consultas rápidas por activo y tiempo
CREATE INDEX IF NOT EXISTS idx_technical_asset_timestamp
ON technical_results (asset_id, timestamp);
