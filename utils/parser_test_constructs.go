package utils

import (
	"fmt"
	t "time"
)

// Const block for verification
const (
	TestConst      = "verified"
	TestConstCount = 100
	EmbeddingTest  = "incremental_v1"
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

// init function for verification
func init() {
	fmt.Println("Initializing parser test constructs")
	TestVar = "initialized"
}
