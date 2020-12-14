import { importFromFile } from "./util";

const cache: { [index: string]: bigint[] } = {};

function main() {
  const instructions = importFromFile("14-input.txt").map(mapToInstruction);

  const memory = insertIntoMemoryPartOne(instructions);

  console.log(
    "Solution Part One: ",
    memory.reduce((a, b) => a + Number(b), 0)
  );

  const memoryPartTwo = insertIntoMemoryPartTwo(instructions);

  let solutionPartTwo = 0;

  for (const val of memoryPartTwo.values()) {
    solutionPartTwo += val;
  }

  console.log("Solution Part Two: ", solutionPartTwo);
}

function insertIntoMemoryPartOne(instructions: Instruction[]): bigint[] {
  const memory: bigint[] = [];
  let mask = 0n;
  let submask = 0n;

  for (const instruction of instructions) {
    if (instruction.action === "mask") {
      mask = BigInt(parseInt(instruction.value.replace(/X/g, "0"), 2));
      submask = BigInt(
        parseInt(instruction.value.replace(/1/g, "0").replace(/X/g, "1"), 2)
      );
    } else {
      memory[instruction.where] = (BigInt(instruction.value) & submask) | mask;
    }
  }

  return memory;
}

function insertIntoMemoryPartTwo(instructions: Instruction[]): Memory {
  const memory: Memory = new Map();
  let mask = "";

  for (const instruction of instructions) {
    if (instruction.action === "mask") {
      mask = instruction.value;
    } else {
      const adressMask = generateMask(mask, instruction.where);

      const adresses: bigint[] = [];
      generateAdresses(adressMask, adresses);

      for (const adress of adresses) {
        memory.set(Number(adress), instruction.value);
      }
    }
  }

  return memory;
}

function generateMask(baseMask: string, adress: number): string {
  const adressBinary = adress.toString(2).padStart(baseMask.length, "0");
  let computedMask = "";

  for (let i = 0; i < baseMask.length; i++) {
    computedMask +=
      baseMask[i] === "X" ? "X" : baseMask[i] === "1" ? "1" : adressBinary[i];
  }

  return computedMask;
}

function generateAdresses(maskStr: string, masks: bigint[]) {
  const xIndex = maskStr.indexOf("X");

  if (xIndex === -1) {
    masks.push(BigInt(parseInt(maskStr, 2)));
    return;
  }

  if (!cache[maskStr]) {
    cache[maskStr] = [];
    generateAdresses(maskStr.replace("X", "0"), cache[maskStr]);
    generateAdresses(maskStr.replace("X", "1"), cache[maskStr]);
  }

  masks.push(...cache[maskStr]);
}

function mapToInstruction(line: string): Instruction {
  const [action, value] = line.split(" = ");
  if (action === "mask") {
    return {
      action,
      value,
    };
  }

  const [, where] = /\[(\d+)\]/.exec(action) || [];

  return {
    action: "mem",
    value: +value,
    where: +where,
  };
}

type Memory = Map<number, number>;

type Instruction = MaskInstruction | MemInstruction;

interface MaskInstruction {
  action: "mask";
  value: string;
}

interface MemInstruction {
  action: "mem";
  value: number;
  where: number;
}

main();
