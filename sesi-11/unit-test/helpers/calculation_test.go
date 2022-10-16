package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// without testify
// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 40 {
// 		t.Error("result should be 40")
// 	}
// 	fmt.Println("test failed")
// }

// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)

// 	if result != 20 {
// 		// panic("result should be 20")
// 		t.FailNow()
// 	}

// 	fmt.Println("test ended")
// }

// with testify
// func TestFailSum(t *testing.T) {
// 	result := Sum(10, 10)
// 	require.Equal(t, 40, result, "result should be 40")
// 	fmt.Println("test failed")
// }

// func TestSum(t *testing.T) {
// 	result := Sum(10, 10)
// 	require.Equal(t, 20, result, "result should be 20")
// 	fmt.Println("test Ended")
// }

// Table Test
func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  Sum(20, 20),
			expected: 40,
			errMsg:   "result should be 40",
		}, {
			request:  Sum(20, 30),
			expected: 50,
			errMsg:   "result should be 50",
		}, {
			request:  Sum(20, 40),
			expected: 60,
			errMsg:   "result should be 60",
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
