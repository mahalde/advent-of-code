package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"math/big"
	"sort"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type Monkey struct {
	items           []*big.Int
	operation       Operation
	testDivisibleBy *big.Int
	throwToOnTrue   int
	throwToOnFalse  int
	itemsInspected  *big.Int
}

func (m *Monkey) MoveAllItems(monkeys []*Monkey, withDecrease bool, divisor *big.Int) {
	for i := range m.items {
		m.itemsInspected.Add(m.itemsInspected, big.NewInt(1))
		m.executeOperation(i, withDecrease, divisor)

		var monkeyToThrowTo *Monkey
		if big.NewInt(0).Mod(m.items[i], m.testDivisibleBy).Int64() == 0 {
			monkeyToThrowTo = monkeys[m.throwToOnTrue]
		} else {
			monkeyToThrowTo = monkeys[m.throwToOnFalse]
		}

		monkeyToThrowTo.items = append(monkeyToThrowTo.items, m.items[i])
	}

	m.items = []*big.Int{}
}

func (m *Monkey) executeOperation(index int, withDecrease bool, divisor *big.Int) {
	var amount *big.Int
	if m.operation.amountIsOld {
		amount = m.items[index]
	} else {
		amount = big.NewInt(m.operation.amount)
	}

	if m.operation.operation == '+' {
		m.items[index].Add(m.items[index], amount)
	} else if m.operation.operation == '*' {
		m.items[index].Mul(m.items[index], amount)
	}

	if withDecrease {
		m.items[index].Div(m.items[index], big.NewInt(3))
	} else {
		// math magic
		m.items[index].Mod(m.items[index], divisor)
	}
}

type Operation struct {
	amount      int64
	amountIsOld bool
	operation   rune
}

func main() {
	input := files.ReadFile(11, 2022, "\n\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) *big.Int {
	return solve(input, 20, true)
}

func solvePart2(input []string) *big.Int {
	return solve(input, 10000, false)
}

func solve(input []string, rounds int, withDecrease bool) *big.Int {
	var monkeys []*Monkey
	divisor := big.NewInt(1)
	for _, block := range input {
		monkey := parseToMonkey(block)
		monkeys = append(monkeys, monkey)
		divisor.Mul(divisor, monkey.testDivisibleBy)
	}

	for round := 0; round < rounds; round++ {
		for _, monkey := range monkeys {
			monkey.MoveAllItems(monkeys, withDecrease, divisor)
		}
	}

	var itemsInspected []*big.Int
	for _, monkey := range monkeys {
		itemsInspected = append(itemsInspected, monkey.itemsInspected)
	}

	sort.Slice(itemsInspected, func(i, j int) bool {
		return itemsInspected[i].Int64() > itemsInspected[j].Int64()
	})
	return big.NewInt(0).Mul(itemsInspected[0], itemsInspected[1])
}

func parseToMonkey(block string) *Monkey {
	lines := strings.Split(block, "\n")
	startingItems := conv.ToBigIntSlice(strings.Split(lines[1][18:], ", "))
	operation := Operation{
		operation: rune(lines[2][23]),
	}

	if lines[2][25:] == "old" {
		operation.amountIsOld = true
	} else {
		operation.amount = int64(conv.ToInt(lines[2][25:]))
	}

	testDivisibleBy := conv.ToInt(lines[3][21:])
	throwToOnTrue := conv.ToInt(lines[4][29:])
	throwToOnFalse := conv.ToInt(lines[5][30:])

	return &Monkey{
		items:           startingItems,
		operation:       operation,
		testDivisibleBy: big.NewInt(int64(testDivisibleBy)),
		throwToOnTrue:   throwToOnTrue,
		throwToOnFalse:  throwToOnFalse,
		itemsInspected:  big.NewInt(0),
	}
}
