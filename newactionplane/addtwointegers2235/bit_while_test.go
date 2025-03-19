package addtwointegers2235

import "testing"

func TestSum2(t *testing.T) {
	tests := []struct {
		num1     int
		num2     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
		{-50, -50, -100},
	}

	for _, test := range tests {
		result := sum2(test.num1, test.num2)
		if result != test.expected {
			t.Errorf("sum2(%d, %d) = %d; expected %d", test.num1, test.num2, result, test.expected)
		}
	}
}
