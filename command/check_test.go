package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersionToCompareGiven(t *testing.T) {
	result := getVersionToCompare("1.0", []string{"0.9", "1.0", "1.1"})
	assert.Equal(t, "1.0", result)
}

func TestGetVersionToCompareLatest(t *testing.T) {
	result := getVersionToCompare("latest", []string{"0.9", "1.0", "1.1"})
	assert.Equal(t, "1.1", result)
}
