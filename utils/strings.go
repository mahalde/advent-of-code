package utils

func Unique(str string) bool {
	m := make(map[rune]bool)
	for _, char := range str {
		_, ok := m[char]
		if ok {
			return false
		}
		m[char] = true
	}

	return true
}
