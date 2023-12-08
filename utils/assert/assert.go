package assert

import "testing"

func IntEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func StringEquals(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
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
