package domain

type Timeframe string

const (
	TimeframeShort  Timeframe = "short"
	TimeframeMedium Timeframe = "medium"
	TimeframeLong   Timeframe = "long"
)

type Trend string

const (
	TrendBullish  Trend = "bullish"
	TrendBearish  Trend = "bearish"
	TrendSideways Trend = "sideways"
)

type Risk string

const (
	RiskLow    Risk = "low"
	RiskMedium Risk = "medium"
	RiskHigh   Risk = "high"
)

type IndicatorSignal struct {
	Indicator string  `json:"indicator"` // e.g. "sma_cross", "stochastic"
	Signal    int     `json:"signal"`    // -1 bearish, 0 neutral, +1 bullish
	Weight    float64 `json:"weight"`    // for future tuning
	Why       string  `json:"why"`       // human explanation
}

type TimeframeAnalysis struct {
	Timeframe  Timeframe         `json:"timeframe"`
	Trend      Trend             `json:"trend"`
	Confidence float64           `json:"confidence"` // 0..100
	Risk       Risk              `json:"risk"`
	Signals    []IndicatorSignal `json:"signals"` // justification list
}

type AssetAnalysisResponse struct {
	Symbol string            `json:"symbol"`
	Short  TimeframeAnalysis `json:"short"`
	Medium TimeframeAnalysis `json:"medium"`
	Long   TimeframeAnalysis `json:"long"`
}
