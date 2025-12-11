INSERT INTO assets (symbol, name, market, currency, exchange, type)
VALUES
    -- USA
    ('AAPL', 'Apple Inc.', 'USA', 'USD', 'NASDAQ', 'stock'),
    ('MSFT', 'Microsoft Corp.', 'USA', 'USD', 'NASDAQ', 'stock'),
    ('AMZN', 'Amazon.com Inc.', 'USA', 'USD', 'NASDAQ', 'stock'),
    ('TSLA', 'Tesla Inc.', 'USA', 'USD', 'NASDAQ', 'stock'),
    ('NVDA', 'NVIDIA Corp.', 'USA', 'USD', 'NASDAQ', 'stock'),

    -- CRYPTO
    ('BTC', 'Bitcoin', 'CRYPTO', 'USD', 'BINANCE', 'crypto'),
    ('ETH', 'Ethereum', 'CRYPTO', 'USD', 'BINANCE', 'crypto'),
    ('SOL', 'Solana', 'CRYPTO', 'USD', 'BINANCE', 'crypto'),
    ('XRP', 'Ripple', 'CRYPTO', 'USD', 'BINANCE', 'crypto'),
    ('ADA', 'Cardano', 'CRYPTO', 'USD', 'BINANCE', 'crypto'),

    -- MERVAL (Argentina)
    ('GGAL', 'Grupo Financiero Galicia', 'MERVAL', 'ARS', 'BYMA', 'stock'),
    ('YPFD', 'YPF S.A.', 'MERVAL', 'ARS', 'BYMA', 'stock'),
    ('PAMP', 'Pampa Energía S.A.', 'MERVAL', 'ARS', 'BYMA', 'stock'),
    ('BMA', 'Banco Macro', 'MERVAL', 'ARS', 'BYMA', 'stock'),
    ('COME', 'Sociedad Comercial del Plata', 'MERVAL', 'ARS', 'BYMA', 'stock');
