package utils

import (
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
