package testutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadingFixtures(t *testing.T) {
	fixtures, err := LoadFixtures("testdata")

	assert.NoError(t, err)
	assert.Len(t, fixtures, 1)
	assert.Contains(t, fixtures, "fixture.yml")
	assert.Contains(t, fixtures["fixture.yml"], "test-fixture")
	assert.Contains(t, fixtures["fixture.yml"], "test-fixture2")
}
