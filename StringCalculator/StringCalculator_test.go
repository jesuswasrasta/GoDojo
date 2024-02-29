package StringCalculator

import "testing"

func Test_text(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"1,2", 3},
	}

	stringCalculator := StringCalculator{}
	for _, test := range tests {
		var result, _ = stringCalculator.add(test.input)
		if result != test.expected {
			t.Errorf("Add(\"%s\") = %d; want %d", test.input, result, test.expected)
		}
	}
}
