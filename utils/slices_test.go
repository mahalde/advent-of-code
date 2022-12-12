package utils

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	got := []string{"hello", "world"}
	want := []string{"world", "hello"}

	ReverseSlice(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		slice []int
		max   int
	}{
		{[]int{2, 100, -4, 300, 20}, 300},
		{[]int{3000000, -20000, 400, 234}, 3000000},
		{[]int{0, 0, 1, 0, 0}, 1},
		{[]int{math.MinInt, math.MaxInt}, math.MaxInt},
	}

	for i, test := range cases {
		t.Run(fmt.Sprint(i+1), func(t *testing.T) {
			got := Max(test.slice)

			AssertIntEquals(t, got, test.max)
		})
	}
}
