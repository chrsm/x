package bd

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestFOAD(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "bad")
}
