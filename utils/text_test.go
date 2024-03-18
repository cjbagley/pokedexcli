package utils

import "testing"

func TestProperTitleCase(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "test",
			expected: "Test",
		},
		{
			input:    "heRe is Another test. does this work?",
			expected: "HeRe is Another test. does this work?",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "The same.",
			expected: "The same.",
		},
	}

	for _, c := range cases {
		actual := UcFirst(c.input)

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("ProperTitleCase(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
