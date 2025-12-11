DELETE FROM assets
WHERE symbol IN (
    'AAPL', 'MSFT', 'AMZN', 'TSLA', 'NVDA',
    'BTC', 'ETH', 'SOL', 'XRP', 'ADA',
    'GGAL', 'YPFD', 'PAMP', 'BMA', 'COME'
);
