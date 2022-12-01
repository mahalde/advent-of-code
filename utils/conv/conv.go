package conv

import "strconv"

func ToIntSlice(slice []string) []int {
	sliceToReturn := []int{}

	for _, current := range slice {
		convertedString, err := strconv.Atoi(current)

		if err != nil {
			panic(err)
		}

		sliceToReturn = append(sliceToReturn, convertedString)
	}

	return sliceToReturn
}

func ToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}
