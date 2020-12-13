import { importFromFile } from "./util";

function main() {
  const input = importFromFile("13-input.txt");
  const earliestDepart = +input[0];
  const availableLines = input[1]
    .split(",")
    .filter((line) => line !== "x")
    .map(Number);
  const allLines = input[1].split(",");

  const [busId, timeToWait] = findEarliestDepartment(
    earliestDepart,
    availableLines
  );

  console.log("Solution Part One: ", busId * timeToWait);

  const timestamp = findTimestampWhereEveryBusDeparts(allLines);

  console.log("Solution Part Two: ", timestamp);
}

function findEarliestDepartment(
  minimumDepartment: number,
  lines: number[]
): [busId: number, timeToWait: number] {
  let earliestDepartment = Infinity;
  let busId = 0;

  for (const line of lines) {
    let firstDepartmentForLine = minimumDepartment;
    while (firstDepartmentForLine % line !== 0) {
      firstDepartmentForLine++;
    }

    if (firstDepartmentForLine < earliestDepartment) {
      earliestDepartment = firstDepartmentForLine;
      busId = line;
    }
  }

  return [busId, earliestDepartment - minimumDepartment];
}

function findTimestampWhereEveryBusDeparts(lines: string[]): number {
  const searchSpace = lines
    .map((lineId, offset) => ({
      lineId,
      cadence: +lineId,
      offset,
    }))
    .filter((line) => line.lineId !== "x")
    .sort((a, b) => a.offset - b.offset);

  let time = 0;
  const firstBus = searchSpace.shift();
  let increment = firstBus?.cadence ?? 0;

  searchSpace.forEach((bus) => {
    let remainder;

    do {
      time = time + increment;
      remainder = (time + bus.offset) % bus.cadence;
    } while (remainder !== 0);
    increment *= bus.cadence;
  });

  return time;
}

main();
