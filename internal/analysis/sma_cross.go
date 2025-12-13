package analysis

import "keisan-aire/internal/domain"

// SMACrossSignal genera una señal por cruce de medias.
// fast > slow  → +1 (alcista)
// fast < slow  → -1 (bajista)
// igual / nil  →  0 (neutral)
func SMACrossSignal(fast, slow *float64, timeframe domain.Timeframe) domain.IndicatorSignal {
	if fast == nil || slow == nil {
		return domain.IndicatorSignal{
			Indicator: "sma_cross",
			Signal:    0,
			Weight:    1.0,
			Why:       "No hay suficientes datos para calcular medias móviles",
		}
	}

	if *fast > *slow {
		return domain.IndicatorSignal{
			Indicator: "sma_cross",
			Signal:    1,
			Weight:    1.0,
			Why:       "La media móvil rápida está por encima de la lenta, indicando tendencia alcista",
		}
	}

	if *fast < *slow {
		return domain.IndicatorSignal{
			Indicator: "sma_cross",
			Signal:    -1,
			Weight:    1.0,
			Why:       "La media móvil rápida está por debajo de la lenta, indicando tendencia bajista",
		}
	}

	return domain.IndicatorSignal{
		Indicator: "sma_cross",
		Signal:    0,
		Weight:    1.0,
		Why:       "Las medias móviles están alineadas, sin una tendencia clara",
	}
}
