CREATE TABLE IF NOT EXISTS statistical_results (
    id SERIAL PRIMARY KEY,

    -- Referencia al activo
    asset_id INTEGER NOT NULL,

    -- Timestamp del cálculo
    timestamp TIMESTAMPTZ NOT NULL,

    -- Estadísticas simples
    mean NUMERIC(18,6),
    stddev NUMERIC(18,6),
    variance NUMERIC(18,6),

    -- Correlaciones (por ahora simple)
    correlation_close_volume NUMERIC(18,6),

    -- Regresión lineal simple
    regression_slope NUMERIC(18,6),
    regression_intercept NUMERIC(18,6),
    regression_r2 NUMERIC(18,6),

    -- ANOVA (single-factor por ahora)
    anova_f NUMERIC(18,6),
    anova_p NUMERIC(18,6),

    -- Relaciones
    CONSTRAINT fk_asset_stats
        FOREIGN KEY (asset_id)
        REFERENCES assets(id)
        ON DELETE CASCADE
);

-- Índices para rapidez
CREATE INDEX IF NOT EXISTS idx_stats_asset_timestamp
ON statistical_results (asset_id, timestamp);
