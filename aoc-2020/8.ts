import { importFromFile } from "./util";

function main() {
  const input = importFromFile("8-input.txt");
  const instructions = input.map(mapToInstrucions);

  findSolutionPartOne(instructions);
  findSolutionPartTwo(instructions);
}

function findSolutionPartOne(instructions: Instruction[]): void {
  const accumulator = executeInstructions(instructions);

  console.log("Solution Part One: ", accumulator);
}

function findSolutionPartTwo(instructions: Instruction[]): void {
  // Deep copy the array
  const newInstructions: Instruction[] = instructions.map(
    ({ operation, argument }) => ({
      operation,
      argument,
    })
  );

  let changedIndex = 0;

  while (true) {
    while (newInstructions[changedIndex].operation === "acc") {
      changedIndex++;
    }

    const instruction = newInstructions[changedIndex];

    instruction.operation = instruction.operation === "nop" ? "jmp" : "nop";

    if (doesTerminate(newInstructions)) {
      break;
    }

    instruction.operation = instructions[changedIndex].operation;
    changedIndex++;
  }

  const accumulator = executeInstructions(newInstructions);

  console.log("Solution Part Two: ", accumulator);
}

function doesTerminate(instructions: Instruction[]): boolean {
  const executedIndexes: number[] = [];
  let nextIndex = 0;

  while (nextIndex < instructions.length) {
    executedIndexes.push(nextIndex);
    const { operation, argument } = instructions[nextIndex];

    if (operation === "jmp") {
      nextIndex += argument;
    } else {
      nextIndex++;
    }

    if (executedIndexes.includes(nextIndex)) {
      return false;
    }
  }

  return true;
}

function executeInstructions(instructions: Instruction[]): number {
  const executedIndexes: number[] = [];
  let accumulator = 0;
  let nextIndex = 0;

  while (
    !executedIndexes.includes(nextIndex) &&
    nextIndex < instructions.length
  ) {
    executedIndexes.push(nextIndex);
    const { operation, argument } = instructions[nextIndex];

    if (operation === "acc") {
      accumulator += argument;
    } else if (operation === "jmp") {
      nextIndex += argument;
      continue;
    }

    nextIndex++;
  }

  return accumulator;
}

function mapToInstrucions(line: string): Instruction {
  const [operation, argument] = line.split(" ");
  return {
    operation: operation as any,
    argument: +argument,
  };
}

interface Instruction {
  operation: "acc" | "jmp" | "nop";
  argument: number;
}

main();
