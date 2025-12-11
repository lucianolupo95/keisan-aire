CREATE TABLE IF NOT EXISTS assets (
    id SERIAL PRIMARY KEY,
    symbol VARCHAR(20) NOT NULL UNIQUE,
    name VARCHAR(100),
    market VARCHAR(50),       -- USA, CRYPTO, MERVAL, JAPAN, CHINA
    currency VARCHAR(10),     -- USD, ARS, JPY, etc.
    exchange VARCHAR(50),     -- NASDAQ, NYSE, BYMA, BINANCE, etc.
    type VARCHAR(50),         -- stock, crypto, index, etf
    created_at TIMESTAMP DEFAULT NOW()
);
