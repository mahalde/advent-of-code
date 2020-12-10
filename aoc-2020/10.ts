import { importFromFile } from "./util";

const joltages = importFromFile("10-input.txt").map(Number);
const deviceJolt = Math.max(...joltages) + 3;
const cache: { [index: number]: number } = {};

function main() {
  const joltageDifferences = getJoltageDifferences();

  console.log(
    "Solution Part One: ",
    joltageDifferences[1] * joltageDifferences[3]
  );

  console.log("Solution Part Two: ", getDistinctArrangements());
}

function getJoltageDifferences(): number[] {
  const differences: number[] = Array(4).fill(0);

  let prevAdapter = Math.min(...joltages);
  let nextAdapter = prevAdapter;

  // Outlet to first adapter
  differences[prevAdapter]++;

  while (nextAdapter !== deviceJolt - 3) {
    prevAdapter = nextAdapter;
    nextAdapter = Math.min(
      ...joltages.filter(
        (joltage) => joltage <= prevAdapter + 3 && joltage > prevAdapter
      )
    );

    differences[nextAdapter - prevAdapter]++;
  }

  // Last adapter to device
  differences[3]++;

  return differences;
}

function getDistinctArrangements(adapter = 0): number {
  if (adapter === deviceJolt - 3) {
    return 1;
  }

  let amount = 0;

  const nextAdapters = joltages.filter(
    (joltage) => joltage <= adapter + 3 && joltage > adapter
  );

  if (!nextAdapters.length) {
    return 0;
  }

  for (const nextAdapter of nextAdapters) {
    if (!cache[nextAdapter]) {
      cache[nextAdapter] = getDistinctArrangements(nextAdapter);
    }
    amount += cache[nextAdapter];
  }

  return amount;
}

main();
