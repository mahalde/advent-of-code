import { importFromFile } from "./util";

function main() {
  const allNumbers = importFromFile("1-input.txt").map(Number);

  const solutionPartOne = getSolutionPartOne(allNumbers);

  const solutionPartTwo = getSolutionPartTwo(allNumbers);

  console.log(solutionPartOne);
  console.log(solutionPartTwo);
}

function getSolutionPartOne(numbers: number[]): number {
  for (const [index, num] of numbers.entries()) {
    const restOfNumbers = numbers.slice(index);

    for (const secondNum of restOfNumbers) {
      if (num + secondNum == 2020) {
        return num * secondNum;
      }
    }
  }

  throw new Error("No two numbers found!");
}

function getSolutionPartTwo(numbers: number[]): number {
  for (const [firstIndex, firstNum] of numbers.entries()) {
    const firstRestOfNumbers = numbers.slice(firstIndex);

    for (const [secondIndex, secondNum] of firstRestOfNumbers.entries()) {
      const secondRestOfNumbers = firstRestOfNumbers.slice(secondIndex);

      for (const thirdNum of secondRestOfNumbers) {
        if (firstNum + secondNum + thirdNum === 2020) {
          return firstNum * secondNum * thirdNum;
        }
      }
    }
  }

  throw new Error("No three numbers found!");
}

main();
