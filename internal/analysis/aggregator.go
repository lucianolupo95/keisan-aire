package analysis

import "keisan-aire/internal/domain"

// AggregateSignals combina señales y decide tendencia, confianza y riesgo.
func AggregateSignals(timeframe domain.Timeframe, signals []domain.IndicatorSignal) domain.TimeframeAnalysis {
	score := 0.0
	totalWeight := 0.0

	for _, s := range signals {
		score += float64(s.Signal) * s.Weight
		totalWeight += s.Weight
	}

	confidence := 0.0
	if totalWeight > 0 {
		confidence = (abs(score) / totalWeight) * 100
	}

	trend := domain.TrendSideways
	if score > 0 {
		trend = domain.TrendBullish
	} else if score < 0 {
		trend = domain.TrendBearish
	}

	risk := domain.RiskHigh
	if confidence >= 70 {
		risk = domain.RiskLow
	} else if confidence >= 40 {
		risk = domain.RiskMedium
	}

	return domain.TimeframeAnalysis{
		Timeframe:  timeframe,
		Trend:      trend,
		Confidence: round(confidence),
		Risk:       risk,
		Signals:    signals,
	}
}

func abs(v float64) float64 {
	if v < 0 {
		return -v
	}
	return v
}

func round(v float64) float64 {
	if v < 0 {
		return 0
	}
	return float64(int(v + 0.5))
}
