import { importFromFile } from "./util";

function main() {
  const input = importFromFile("9-input.txt").map(Number);

  const invalidNumber = findInvalidNumber(input, 25);

  console.log("Solution Part One: ", invalidNumber);

  const setOfNumbers = findNumbersWhichAddToInvalidNumber(input, invalidNumber);

  const encryptionWeakness =
    Math.min(...setOfNumbers) + Math.max(...setOfNumbers);

  console.log("Solution Part Two: ", encryptionWeakness);
}

function findInvalidNumber(numbers: number[], preamble: number): number {
  for (let i = 0; i < numbers.length; i++) {
    if (i - preamble < 0) continue;

    const numbersToAdd = numbers.slice(i - preamble, i);
    let isAddable = false;

    for (const [index, value] of numbersToAdd.entries()) {
      for (const secondValue of numbersToAdd.slice(index + 1)) {
        if (value + secondValue === numbers[i]) {
          isAddable = true;
          break;
        }
      }

      if (isAddable) break;
    }

    if (!isAddable) {
      return numbers[i];
    }
  }
  throw new Error("No invalid number found");
}

function findNumbersWhichAddToInvalidNumber(
  numbers: number[],
  invalidNumber: number
): number[] {
  let addedNumbers: number[] = [];

  for (let i = 0; i < numbers.length; i++) {
    addedNumbers.push(numbers[i]);

    for (let j = i + 1; j < numbers.length; j++) {
      addedNumbers.push(numbers[j]);

      if (addNumbers(addedNumbers) > invalidNumber) {
        addedNumbers = [];
        break;
      }

      if (addNumbers(addedNumbers) === invalidNumber) {
        return addedNumbers;
      }
    }
  }

  throw new Error(`No numbers found which add to ${invalidNumber}`);
}

function addNumbers(numbers: number[]): number {
  return numbers.reduce((a, b) => a + b, 0);
}

main();
