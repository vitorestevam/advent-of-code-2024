package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		expected := 18
		output := Part1("input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("real input", func(t *testing.T) {
		expected := 2536
		output := Part1("input.txt")
		assert.Equal(t, expected, output)
	})
}

func TestPart2(t *testing.T) {
	t.Run("test input", func(t *testing.T) {
		expected := 14
		output := Part2("input_test.txt")
		assert.Equal(t, expected, output)
	})
	t.Run("real input", func(t *testing.T) {
		expected := 1875
		output := Part2("input.txt")
		assert.Equal(t, expected, output)
	})
}
