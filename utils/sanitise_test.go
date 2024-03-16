package utils

import "testing"

func TestSanitisePromptInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "test",
			expected: []string{"test"},
		},
		{
			input:    "Test",
			expected: []string{"test"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "   HelLo woRld  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "   !this-is-okAy!  ",
			expected: []string{"this-is-okay"},
		},
		{
			input:    `|Symbols*should!be!#'?!\"^$*()removed`,
			expected: []string{"symbols", "should", "be", "removed"},
		},
	}

	for _, c := range cases {
		actual := SanitisePromptInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("SanitisePrompInput lengths different: expected %v, got %v", c.expected, actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("SanitisePromptInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
