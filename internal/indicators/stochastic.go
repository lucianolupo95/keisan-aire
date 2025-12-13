package indicators

// StochasticK calcula el %K del estocástico.
// prices: precios de cierre
// window: período (ej. 14)
// Devuelve nil si no hay datos suficientes.
func StochasticK(prices []float64, window int) *float64 {
	if len(prices) < window || window <= 0 {
		return nil
	}

	low := prices[len(prices)-window]
	high := prices[len(prices)-window]

	for i := len(prices) - window; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}
		if prices[i] > high {
			high = prices[i]
		}
	}

	if high == low {
		return nil
	}

	k := (prices[len(prices)-1] - low) / (high - low) * 100
	return &k
}
