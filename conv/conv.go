package conv

import (
	"math/big"
	"strconv"
)

func ToIntSlice(slice []string) []int {
	var sliceToReturn []int

	for _, current := range slice {
		sliceToReturn = append(sliceToReturn, ToInt(current))
	}

	return sliceToReturn
}

func ToBigIntSlice(slice []string) []*big.Int {
	var sliceToReturn []*big.Int

	for _, current := range slice {
		sliceToReturn = append(sliceToReturn, big.NewInt(int64(ToInt(current))))
	}

	return sliceToReturn
}

func ToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return num
}
