import { cloneDeep, isEqual } from "lodash";
import { importFromFile } from "./util";

function main() {
  const input = importFromFile("11-input.txt").map((line) => line.split(""));

  const solutionPartOne = getAllOccupiedSeats(
    input,
    countNearbyOccupiedSeats,
    4
  );

  console.log("Solution Part One: ", solutionPartOne);

  const solutionPartTwo = getAllOccupiedSeats(
    input,
    countOccupiedSeatsIn8Directions,
    5
  );

  console.log("Solution Part Two: ", solutionPartTwo);
}

function getAllOccupiedSeats(
  input: string[][],
  countFn: Function,
  numberToBecomeEmpty: number
): number {
  let seats: string[][] = [];
  let nextSeats: string[][] = cloneDeep(input);

  do {
    seats = cloneDeep(nextSeats);
    for (const [lineIndex, line] of seats.entries()) {
      nextSeats[lineIndex] = [];
      for (const [seatIndex, seat] of line.entries()) {
        if (seat === ".") {
          nextSeats[lineIndex][seatIndex] = ".";
        } else if (seat === "L" && countFn(seats, lineIndex, seatIndex) === 0) {
          nextSeats[lineIndex][seatIndex] = "#";
        } else if (
          seat === "#" &&
          countFn(seats, lineIndex, seatIndex) >= numberToBecomeEmpty
        ) {
          nextSeats[lineIndex][seatIndex] = "L";
        } else {
          nextSeats[lineIndex][seatIndex] = seats[lineIndex][seatIndex];
        }
      }
    }
  } while (!isEqual(seats, nextSeats));

  return countAllOccupiedSeats(seats);
}

function countNearbyOccupiedSeats(
  seats: string[][],
  lineIndex: number,
  seatIndex: number
) {
  let count = 0;
  for (let l = lineIndex - 1; l <= lineIndex + 1; l++) {
    for (let s = seatIndex - 1; s <= seatIndex + 1; s++) {
      if (s === seatIndex && l === lineIndex) continue;
      if (seats[l]?.[s] === "#") count++;
    }
  }
  return count;
}

function countOccupiedSeatsIn8Directions(
  seats: string[][],
  lineIndex: number,
  seatIndex: number
): number {
  let count = 0;

  // Vertical up
  for (let l = lineIndex - 1; l >= 0; l--) {
    if (seats[l][seatIndex] === "L") break;
    if (seats[l][seatIndex] === "#") {
      count++;
      break;
    }
  }

  // Vertical down
  for (let l = lineIndex + 1; l < seats.length; l++) {
    if (seats[l][seatIndex] === "L") break;
    if (seats[l][seatIndex] === "#") {
      count++;
      break;
    }
  }

  // Horizontal left
  for (let s = seatIndex - 1; s >= 0; s--) {
    if (seats[lineIndex][s] === "L") break;
    if (seats[lineIndex][s] === "#") {
      count++;
      break;
    }
  }

  // Horizontal right
  for (let s = seatIndex + 1; s < seats[lineIndex].length; s++) {
    if (seats[lineIndex][s] === "L") break;
    if (seats[lineIndex][s] === "#") {
      count++;
      break;
    }
  }

  // Diagonal up left
  for (let l = lineIndex - 1, s = seatIndex - 1; l >= 0 && s >= 0; l--, s--) {
    if (seats[l][s] === "L") break;
    if (seats[l][s] === "#") {
      count++;
      break;
    }
  }

  // Diagonal up right
  for (
    let l = lineIndex - 1, s = seatIndex + 1;
    l >= 0 && s < seats[lineIndex].length;
    l--, s++
  ) {
    if (seats[l][s] === "L") break;
    if (seats[l][s] === "#") {
      count++;
      break;
    }
  }

  // Diagonal down left
  for (
    let l = lineIndex + 1, s = seatIndex - 1;
    l < seats.length && s >= 0;
    l++, s--
  ) {
    if (seats[l][s] === "L") break;
    if (seats[l][s] === "#") {
      count++;
      break;
    }
  }

  // Diagonal down right
  for (
    let l = lineIndex + 1, s = seatIndex + 1;
    l < seats.length && s < seats[lineIndex].length;
    l++, s++
  ) {
    if (seats[l][s] === "L") break;
    if (seats[l][s] === "#") {
      count++;
      break;
    }
  }

  return count;
}

function countAllOccupiedSeats(seats: string[][]) {
  let count = 0;
  for (const line of seats) {
    for (const seat of line) {
      if (seat === "#") count++;
    }
  }

  return count;
}

main();
