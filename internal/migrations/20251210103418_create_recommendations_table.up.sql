CREATE TABLE IF NOT EXISTS recommendations (
    id SERIAL PRIMARY KEY,

    -- Relación con el activo
    asset_id INTEGER NOT NULL,

    -- Timestamp de la recomendación
    timestamp TIMESTAMPTZ NOT NULL,

    -- Acción recomendada
    action VARCHAR(10) NOT NULL,     -- "buy", "sell", "hold"

    -- Nivel de riesgo percibido por el sistema
    risk_level VARCHAR(20),          -- "low", "medium", "high"

    -- Motivación general
    reason TEXT,

    -- Señales utilizadas (técnico, estadístico, ML, fundamental)
    technical_signal NUMERIC(18,6),
    statistical_signal NUMERIC(18,6),
    ml_signal NUMERIC(18,6),
    fundamental_signal NUMERIC(18,6),

    -- JSON con detalles extendidos
    details JSONB,

    CONSTRAINT fk_asset_reco
        FOREIGN KEY (asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

-- Índice para obtener rápidamente la última recomendación por activo
CREATE INDEX IF NOT EXISTS idx_reco_asset_timestamp
ON recommendations (asset_id, timestamp DESC);
