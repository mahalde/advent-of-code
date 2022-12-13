package main

import (
	_ "embed"
	"github.com/mahalde/advent-of-code/utils/files"
	"math/big"
	"reflect"
	"testing"
)

var (
	//go:embed testdata/data
	file string

	input = files.ParseFile(file, "\n\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input)
	if solution.Int64() != 10605 {
		t.Errorf("got %v, want %v", solution, 10605)
	}
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)
	want := big.NewInt(2713310158)

	if solution.Cmp(want) != 0 {
		t.Errorf("got %v, want %v", solution, want)
	}
}

func TestParseToMonkey(t *testing.T) {
	got := parseToMonkey(input[0])
	want := &Monkey{
		items: []*big.Int{big.NewInt(79), big.NewInt(98)},
		operation: Operation{
			amount:    19,
			operation: '*',
		},
		testDivisibleBy: big.NewInt(23),
		throwToOnTrue:   2,
		throwToOnFalse:  3,
		itemsInspected:  big.NewInt(0),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	got = parseToMonkey(input[2])
	want = &Monkey{
		items: []*big.Int{big.NewInt(79), big.NewInt(60), big.NewInt(97)},
		operation: Operation{
			operation:   '*',
			amountIsOld: true,
		},
		testDivisibleBy: big.NewInt(13),
		throwToOnTrue:   1,
		throwToOnFalse:  3,
		itemsInspected:  big.NewInt(0),
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestMonkey_MoveAllItems(t *testing.T) {
	var got []*Monkey
	for _, block := range input {
		got = append(got, parseToMonkey(block))
	}

	var want []*Monkey
	for _, block := range input {
		want = append(want, parseToMonkey(block))
	}

	want[0].items = []*big.Int{}
	want[0].itemsInspected = big.NewInt(2)
	want[3].items = append(want[3].items, big.NewInt(500), big.NewInt(620))

	got[0].MoveAllItems(got, true, big.NewInt(0))

	if !reflect.DeepEqual(got[0], want[0]) {
		t.Errorf("first monkey: got %+v, want %+v", got[0], want[0])
	}

	if !reflect.DeepEqual(got[3], want[3]) {
		t.Errorf("second monkey: got %+v, want %+v", got[3], want[3])
	}
}
