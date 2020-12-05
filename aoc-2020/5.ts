import { importFromFile } from "./util";

function main() {
  const boardingPasses = importFromFile("5-input.txt");
  const seats = boardingPasses.map(getSeat);
  const seatIDs = seats.map(([row, column]) => row * 8 + column);
  const maxSeatID = Math.max(...seatIDs);
  const missingSeatID = getMissingSeatID(seatIDs, maxSeatID);

  console.log("Solution Part One: ", maxSeatID);
  console.log("Solution Part Two: ", missingSeatID);
}

function getSeat(pass: string): [row: number, column: number] {
  const mapToLowerOrUpper = (char: string) =>
    char === "F" || char === "L" ? "lower" : "upper";
  const firstSteps = pass.slice(0, 7).split("").map(mapToLowerOrUpper);
  const secondSteps = pass.slice(-3).split("").map(mapToLowerOrUpper);

  const row = binarySpacePartitioning(0, 127, firstSteps);
  const column = binarySpacePartitioning(0, 7, secondSteps);
  return [row, column];
}

function getMissingSeatID(ids: number[], maxID: number): number {
  let missingIDs: number[] = [];

  for (let i = 0; i < maxID; i++) {
    if (!ids.includes(i)) {
      missingIDs.push(i);
    }
  }

  missingIDs = missingIDs.map((id, _, ids) =>
    ids.includes(id + 1) || ids.includes(id - 1) ? 0 : id
  );

  return missingIDs.filter((id) => !!id)[0];
}

function binarySpacePartitioning(
  lower: number,
  upper: number,
  steps: ("lower" | "upper")[],
  index = 0
): number {
  if (lower === upper) {
    return lower;
  }

  if (!steps[index]) {
    throw new Error(`Undefined for index ${index} for steps: ${steps}`);
  }

  const step = steps[index];

  if (step === "lower") {
    upper = lower + Math.floor((upper - lower) / 2);
  } else {
    lower += Math.ceil((upper - lower) / 2);
  }

  return binarySpacePartitioning(lower, upper, steps, index + 1);
}

main();
