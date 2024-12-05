package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		expected := 161
		output := Part1("input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("real input", func(t *testing.T) {
		expected := 178538786
		output := Part1("input.txt")
		assert.Equal(t, expected, output)
	})
}
