package utils

import (
	"math"
	"sort"
)

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func Max(slice []int) int {
	max := math.MinInt
	for _, el := range slice {
		if el > max {
			max = el
		}
	}

	return max
}
