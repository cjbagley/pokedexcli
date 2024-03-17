package utils

import (
	"testing"
)

func TestPassesRandomThreshholdCheck(t *testing.T) {
	cases := []struct {
		input             RandomThresholdCheck
		expected          bool
		seedGivesResultOf int
	}{
		{
			input: RandomThresholdCheck{
				Value: 120,
				Seed:  1,
			},
			seedGivesResultOf: 281,
			expected:          true,
		},
		{
			input: RandomThresholdCheck{
				Value: 255,
				Seed:  10,
			},
			seedGivesResultOf: 254,
			expected:          false,
		},
		{
			input: RandomThresholdCheck{
				Value: 254,
				Seed:  10,
			},
			seedGivesResultOf: 254,
			expected:          true,
		},
		{
			input: RandomThresholdCheck{
				Value: 200,
				Seed:  9,
			},
			seedGivesResultOf: 101,
			expected:          false,
		},
	}

	for _, c := range cases {
		actual := PassesRandomThreshholdCheck(c.input)

		if actual != c.expected {
			t.Errorf("passesRandomThreshholdCheck(%v) == %v, expected %v. Seed should give: %v",
				c.input.Value,
				actual,
				c.expected,
				c.seedGivesResultOf,
			)
		}
	}
}
