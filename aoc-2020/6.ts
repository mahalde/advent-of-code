import { importFromFile } from "./util";

function main() {
  const input = importFromFile("6-input.txt");

  getSolutionPartOne(input);
  getSolutionPartTwo(input);
}

function getSolutionPartOne(input: string[]) {
  const answersInGroup = mapToAnswersInGroup(input);
  const solution = answersInGroup.reduce(countAnswers, 0);

  console.log("Solution Part One: ", solution);
}

function getSolutionPartTwo(input: string[]) {
  const answersPerGroup = mapToAnswersPerGroup(input);
  const answersWithAgreements = answersPerGroup.map(mapToAllAgreements);
  const solution = answersWithAgreements.reduce(countAnswers, 0);

  console.log("Solution Part Two: ", solution);
}

function mapToAnswersInGroup(input: string[]): Set<string>[] {
  const answers: Set<string>[] = [];
  let tempString = "";

  for (const line of input) {
    if (!!line) {
      tempString += line;
    } else {
      tempString.trim();
      answers.push(new Set(tempString));
      tempString = "";
    }
  }

  return answers;
}

function mapToAnswersPerGroup(input: string[]): string[][] {
  const answers: string[][] = [];
  let answer: string[] = [];

  for (const line of input) {
    if (!!line) {
      answer.push(line);
    } else {
      answers.push(answer);
      answer = [];
    }
  }

  return answers;
}

function mapToAllAgreements(groupAnswers: string[]): Set<string> {
  const agreements: Map<string, number> = new Map();

  for (const answer of groupAnswers) {
    for (const singleAnswer of answer) {
      const newNumber = agreements.has(singleAnswer)
        ? (agreements.get(singleAnswer) ?? 1) + 1
        : 1;
      agreements.set(singleAnswer, newNumber);
    }
  }

  const allAgreements: string[] = [];
  for (const [key, value] of agreements) {
    if (value === groupAnswers.length) {
      allAgreements.push(key);
    }
  }

  return new Set(allAgreements);
}

function countAnswers(prev: number, current: Set<string>): number {
  return prev + current.size;
}

main();
