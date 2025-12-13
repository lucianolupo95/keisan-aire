package analysis

import "keisan-aire/internal/domain"

func StochasticSignal(k *float64) domain.IndicatorSignal {
	if k == nil {
		return domain.IndicatorSignal{
			Indicator: "stochastic",
			Signal:    0,
			Weight:    0.5,
			Why:       "No hay suficientes datos para calcular el estocástico",
		}
	}

	if *k < 20 {
		return domain.IndicatorSignal{
			Indicator: "stochastic",
			Signal:    1,
			Weight:    0.5,
			Why:       "El estocástico indica sobreventa, posible rebote alcista",
		}
	}

	if *k > 80 {
		return domain.IndicatorSignal{
			Indicator: "stochastic",
			Signal:    -1,
			Weight:    0.5,
			Why:       "El estocástico indica sobrecompra, posible corrección bajista",
		}
	}

	return domain.IndicatorSignal{
		Indicator: "stochastic",
		Signal:    0,
		Weight:    0.5,
		Why:       "El estocástico no indica condiciones extremas",
	}
}
