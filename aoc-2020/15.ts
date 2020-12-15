import { importFromFile } from "./util";

function main() {
  const startingNumbers = importFromFile("15-input.txt")[0]
    .split(",")
    .map(Number);
  const numOfTurnsPartOne = 2020;
  const numOfTurnsPartTwo = 30_000_000;

  console.log(
    "Solution Part One: ",
    findIndex(playGame(startingNumbers, numOfTurnsPartOne), numOfTurnsPartOne)
  );

  console.log(
    "Solution Part Two: ",
    findIndex(playGame(startingNumbers, numOfTurnsPartTwo), numOfTurnsPartTwo)
  );
}

function playGame(startingNumbers: number[], turns: number): Game {
  const spokenNumbers: Game = new Map();

  for (const [index, num] of startingNumbers.entries()) {
    spokenNumbers.set(num, { lastIndex: index + 1 });
  }

  let lastNumber = spokenNumbers.get(
    startingNumbers[startingNumbers.length - 1]
  );

  for (let i = startingNumbers.length + 1; i <= turns; i++) {
    let key: number;
    if (!!lastNumber?.firstIndex) {
      key = lastNumber.lastIndex - lastNumber.firstIndex;
    } else {
      key = 0;
    }

    lastNumber = {
      firstIndex: spokenNumbers.get(key)?.lastIndex,
      lastIndex: i,
    };
    spokenNumbers.set(key, lastNumber);
  }
  return spokenNumbers;
}

function findIndex(game: Game, index: number): number {
  for (const [key, value] of game.entries()) {
    if (value.lastIndex === index) return key;
  }

  throw new Error(`No value found at index ${index}`);
}

type Game = Map<number, { firstIndex?: number; lastIndex: number }>;

main();
