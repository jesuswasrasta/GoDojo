package StringCalculator

import "testing"

func Test_add_number_1_return_1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
	}

	for _, test := range tests {
		var result, _ = add(test.input)
		if result != test.expected {
			t.Errorf("Add(\"%s\") = %d; want %d", test.input, result, test.expected)
		}
	}
}
