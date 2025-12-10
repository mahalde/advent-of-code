package math

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Delta returns the delta of x.
// Cases:
//
// x is positive, returns 1
//
// x is negative, returns -1
//
// x is 0, returns 0
func Delta(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
