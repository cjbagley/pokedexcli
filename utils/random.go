package utils

import (
	"math/rand"
)

const RANDOM_THRESHOLD_CHECK_DEFAULT_MAX = 600

type RandomThresholdCheck struct {
	Value int
	Max   int
	Seed  int
}

func PassesRandomThreshholdCheck(check RandomThresholdCheck) bool {
	if check.Max == 0 {
		check.Max = RANDOM_THRESHOLD_CHECK_DEFAULT_MAX
	}

	val := 0
	if check.Seed != 0 {
		r := rand.New(rand.NewSource(int64(check.Seed)))
		val = r.Intn(check.Max)
	} else {
		val = rand.Intn(check.Max)
	}

	return check.Value <= val
}
