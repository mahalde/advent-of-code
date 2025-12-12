package assert

import "testing"

func Equals[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func True(t testing.TB, got bool) {
	t.Helper()

	if got == false {
		t.Errorf("wanted %v to be true", got)
	}
}

func False(t testing.TB, got bool) {
	t.Helper()

	if got == true {
		t.Errorf("wanted %v to be false", got)
	}
}

func SlicesEqual[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	Equals(t, len(got), len(want))
	for i, v := range got {
		Equals(t, v, want[i])
	}
}
