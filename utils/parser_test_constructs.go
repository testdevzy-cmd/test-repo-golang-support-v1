package utils

import (
	"fmt"
	"strings"
	t "time"
)

// Const block for verification
const (
	TestConst      = "verified"
	TestConstCount = 102
	EmbeddingTest  = "incremental_v3"
)

// Var block for verification
var (
	TestVar     = "variable"
	TestVarTime = t.Now()
)

// BaseStruct to be embedded
type BaseStruct struct {
	ID        string
	CreatedAt t.Time
}

// EmbedStruct embedding BaseStruct
type EmbedStruct struct {
	BaseStruct
	Name        string
	Description string
}

// ReviewSummaryRow models a row in the review summary table output.
type ReviewSummaryRow struct {
	BaseStruct
	FilePath   string
	Line       int
	Severity   string
	Message    string
	Category   string
	Confidence float64
}

func (r ReviewSummaryRow) DisplayLine() string {
	return fmt.Sprintf("[%s] %s:%d %s", r.Severity, r.FilePath, r.Line, r.Message)
}

func (r ReviewSummaryRow) IsHighPriority() bool {
	return r.Severity == SeverityCritical || r.Severity == SeverityHigh
}

// ReviewSummaryTable aggregates rows for summary rendering tests.
type ReviewSummaryTable struct {
	Rows    []ReviewSummaryRow
	Version string
}

func NewReviewSummaryTable(version string) *ReviewSummaryTable {
	return &ReviewSummaryTable{
		Rows:    make([]ReviewSummaryRow, 0),
		Version: version,
	}
}

func (t *ReviewSummaryTable) AddRow(row ReviewSummaryRow) {
	t.Rows = append(t.Rows, row)
}

func (t *ReviewSummaryTable) CountBySeverity() map[string]int {
	counts := make(map[string]int)
	for _, row := range t.Rows {
		counts[row.Severity]++
	}
	return counts
}

func (t *ReviewSummaryTable) HighPriorityRows() []ReviewSummaryRow {
	high := make([]ReviewSummaryRow, 0)
	for _, row := range t.Rows {
		if row.IsHighPriority() {
			high = append(high, row)
		}
	}
	return high
}

func (t *ReviewSummaryTable) RenderMarkdown() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("## Review Summary (%s)\n\n", t.Version))
	b.WriteString("| Severity | File | Line | Message |\n")
	b.WriteString("|----------|------|------|----------|\n")
	for _, row := range t.Rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %s |\n",
			row.Severity, row.FilePath, row.Line, row.Message))
	}
	return b.String()
}

// init function for verification
func init() {
	fmt.Println("Initializing parser test constructs")
	TestVar = "initialized"
}
