package utils

const MaxIterations = 1000
const DefaultTimeout = 30
const Pi = 3.14159

const (
	ReviewSummaryVersion = "v2"
	SummaryTableRows     = 10
	MinScoreThreshold    = 0.5
	MaxCommentLength     = 4096
)

const (
	SeverityCritical = "critical"
	SeverityHigh     = "high"
	SeverityMedium   = "medium"
	SeverityLow      = "low"
)

var SeverityOrder = map[string]int{
	SeverityCritical: 4,
	SeverityHigh:     3,
	SeverityMedium:   2,
	SeverityLow:      1,
}

func DefaultSeverityWeights() map[string]float64 {
	return map[string]float64{
		SeverityCritical: 1.0,
		SeverityHigh:     0.75,
		SeverityMedium:   0.5,
		SeverityLow:      0.25,
	}
}
