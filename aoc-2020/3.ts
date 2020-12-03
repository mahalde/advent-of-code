import { importFromFile } from "./util";

function main() {
  const inputLines = importFromFile("3-input.txt");

  const inputMatrix = inputLines.map((line) => line.split(""));

  getSolutionPartOne(inputMatrix);
  getSolutionPartTwo(inputMatrix);
}

function getSolutionPartOne(inputMatrix: string[][]) {
  const solution = findTreeEncounters(inputMatrix, 3, 1);

  console.log("Solution Part One: ", solution);
}

function getSolutionPartTwo(inputMatrix: string[][]) {
  const slopes = [
    [1, 1],
    [3, 1],
    [5, 1],
    [7, 1],
    [1, 2],
  ];

  const encounters = slopes.map((slope) =>
    findTreeEncounters(inputMatrix, slope[0], slope[1])
  );

  const solution = encounters.reduce((a, b) => a * b, 1);

  console.log("Solution Part Two: ", solution);
}

function findTreeEncounters(
  inputMatrix: string[][],
  xStep: number,
  yStep: number
): number {
  let x = 0;
  let y = 0;
  let encounters = 0;
  while (y < inputMatrix.length) {
    const line = inputMatrix[y];
    x %= line.length;

    if (line[x] === "#") {
      encounters++;
    }
    x += xStep;
    y += yStep;
  }
  return encounters;
}

main();
