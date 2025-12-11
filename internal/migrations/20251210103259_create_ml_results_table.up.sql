CREATE TABLE IF NOT EXISTS ml_results (
    id SERIAL PRIMARY KEY,

    -- Relación con activo
    asset_id INTEGER NOT NULL,

    -- Timestamp del cálculo/predicción
    timestamp TIMESTAMPTZ NOT NULL,

    -- Tipo de modelo (rf, lr, mlp, lstm, etc.)
    model_type VARCHAR(50),

    -- Predicciones principales
    predicted_direction VARCHAR(10),          -- "up" / "down" / "flat"
    probability_up NUMERIC(18,6),
    probability_down NUMERIC(18,6),

    -- Métricas del modelo
    accuracy NUMERIC(18,6),
    loss NUMERIC(18,6),
    rmse NUMERIC(18,6),

    -- Features importantes (opcional)
    feature_importance JSONB,

    CONSTRAINT fk_asset_ml
        FOREIGN KEY (asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

-- Índice para consultar un activo rápidamente
CREATE INDEX IF NOT EXISTS idx_ml_asset_timestamp
ON ml_results (asset_id, timestamp);
