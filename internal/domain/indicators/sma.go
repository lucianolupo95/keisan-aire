package indicators

// SMA calcula la media móvil simple de una serie de precios.
// window = cantidad de puntos a promediar.
// Devuelve nil si no hay datos suficientes.
func SMA(prices []float64, window int) *float64 {
	if len(prices) < window || window <= 0 {
		return nil
	}

	sum := 0.0
	for i := len(prices) - window; i < len(prices); i++ {
		sum += prices[i]
	}

	avg := sum / float64(window)
	return &avg
}
