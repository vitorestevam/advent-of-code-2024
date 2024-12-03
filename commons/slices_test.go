package commons

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Map UpperCase", func(t *testing.T) {
		entry := []string{"abc", "ABC", "123efg"}
		expected := []string{"ABC", "ABC", "123EFG"}

		result := Map(entry, strings.ToUpper)

		assert.Equal(t, expected, result)
	})
}

func TestReduce(t *testing.T) {
	t.Run("Sum ints", func(t *testing.T) {
		entry := []int{1, 2, 3, 4, 5}
		expected := 15

		fn := func(a, b int) int {
			return a + b
		}

		result := Reduce(entry, fn)

		assert.Equal(t, expected, result)
	})

	t.Run("Concatenate strings", func(t *testing.T) {
		entry := []string{"a", "b", "c", "d", "e"}
		expect := "abcde"

		fn := func(a, b string) string {
			return a + b
		}

		result := Reduce(entry, fn)

		assert.Equal(t, expect, result)
	})
}
