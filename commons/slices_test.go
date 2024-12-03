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
