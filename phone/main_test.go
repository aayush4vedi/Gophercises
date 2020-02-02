package main

import "testing"

// Table Driven Testing`
func TestNormalize(t *testing.T) {
	testCase := []struct {
		input string
		want  string
	}{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"123-456-7890", "1234567890"},
		{"1234567892", "1234567892"},
		{"(123)456-7892", "1234567892"},
	}
	for _, tc := range testCase {
		t.Run(tc.input, func(t *testing.T) {
			actual := tc.want
			got := normalize(tc.input)
			if actual != got {
				t.Errorf("Found %s, wanted %s", got, actual)
			}
		})
	}
}
