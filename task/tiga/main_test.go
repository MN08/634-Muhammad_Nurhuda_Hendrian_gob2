// with testify
package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsNotPalindorme(t *testing.T) {
	result := Sum(10, 10)
	require.Equal(t, 40, result, "result should be 40")
	fmt.Println("test failed")
}

func TestIsPalindorme(t *testing.T) {
	result := Sum(10, 10)
	assert.Equal(t, 20, result, "result should be 20")
	fmt.Println("test Ended")
}
