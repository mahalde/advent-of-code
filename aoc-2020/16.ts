import { importFromFile } from "./util";

function main() {
  const [rules, myTicket, nearbyTickets] = mapInput(
    importFromFile("16-input.txt")
  );

  const solutionPartOne = getSolutionPartOne(rules, nearbyTickets);

  console.log("Solution Part One: ", solutionPartOne);

  const solutionPartTwo = getSolutionPartTwo(rules, myTicket, nearbyTickets);

  console.log("Solution Part Two: ", solutionPartTwo);
}

function getSolutionPartOne(rules: Rule[], tickets: number[][]) {
  const range = new Set(
    rules.reduce((all: number[], rule) => all.concat(...rule.values), [])
  );

  const invalidNumbers: number[] = [];

  for (const ticket of tickets) {
    invalidNumbers.push(...ticket.filter((num) => !range.has(num)));
  }

  return invalidNumbers.reduce((a, b) => a + b, 0);
}

function getSolutionPartTwo(
  rules: Rule[],
  myTicket: number[],
  tickets: number[][]
): number {
  const range = new Set(
    rules.reduce((all: number[], rule) => all.concat(...rule.values), [])
  );

  const validIndexesMap: Map<string, number[]> = new Map();

  const validTickets = tickets.filter((ticket) =>
    ticket.every((num) => range.has(num))
  );

  for (const rule of rules) {
    const indexes = getIndexes(rule.values, validTickets);

    validIndexesMap.set(rule.name, indexes);
  }

  const sortedValidIndexesMap = new Map(
    Array.from(validIndexesMap).sort((a, b) => a[1].length - b[1].length)
  );

  // Key: rule name, value: index;
  const ruleMap: Map<string, number> = new Map();
  const alreadyAssignedIndexes: number[] = [];
  for (const [key, values] of sortedValidIndexesMap) {
    const index = values.filter(
      (num) => !alreadyAssignedIndexes.includes(num)
    )[0];
    ruleMap.set(key, index);
    alreadyAssignedIndexes.push(index);
  }

  let solution = 1;

  for (const [key, index] of ruleMap) {
    if (key.startsWith("departure")) {
      solution *= myTicket[index];
    }
  }

  return solution;
}

function getIndexes(values: number[], tickets: number[][]): number[] {
  const validIndexes: number[] = [];

  loop: for (let i = 0; i < tickets[0].length; i++) {
    for (const ticket of tickets) {
      if (!values.includes(ticket[i])) {
        continue loop;
      }
    }

    validIndexes.push(i);
  }

  return validIndexes;
}

function mapInput(
  input: string[]
): [rules: Rule[], myTicket: number[], nearbyTickets: number[][]] {
  let isMyTicket = false;
  let isRules = true;

  const rules: Rule[] = [];
  let myTicket: number[] = [];
  const nearbyTickets: number[][] = [];

  for (let i = 0; i < input.length; i++) {
    const line = input[i];
    if (isRules) {
      if (!line) {
        isRules = false;
        isMyTicket = true;
        i++;
        continue;
      }
      const values: number[] = [];
      const [name, ranges] = line.split(": ");
      ranges.split(" or ").forEach((range) => {
        const [lower, upper] = range.split("-").map(Number);
        for (let j = lower; j <= upper; j++) values.push(j);
      });

      rules.push({ name, values });
    } else if (isMyTicket) {
      myTicket = line.split(",").map(Number);
      isMyTicket = false;
      i += 2;
    } else {
      nearbyTickets.push(line.split(",").map(Number));
    }
  }

  return [rules, myTicket, nearbyTickets];
}

interface Rule {
  name: string;
  values: number[];
}

main();
