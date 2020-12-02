import { importFromFile } from "./util";

function main() {
  const file = importFromFile("2-input.txt");

  const passwordPolicies = file.map(getInformationFromLine);

  getSolutionPartOne(passwordPolicies);
  getSolutionPartTwo(passwordPolicies);
}

function getSolutionPartOne(passwordPolicies: PasswordPolicy[]) {
  const validPolicies = passwordPolicies.map(isValidPolicyPartOne);

  const solution = validPolicies.reduce((a, b) => +a + +b, 0);

  console.log(solution);
}

function getSolutionPartTwo(passwordPolicies: PasswordPolicy[]) {
  const validPolicies = passwordPolicies.map(isValidPolicyPartTwo);
  console.log(validPolicies);
  const solution = validPolicies.reduce((a, b) => +a + +b, 0);

  console.log(solution);
}

function getInformationFromLine(line: string): PasswordPolicy {
  const [firstPartOfLine, password] = line.split(": ");

  const firstPartRegEx = /(\d{1,2})-(\d{1,2}) (\w)/;

  const [, lowestAmount, highestAmount, givenLetter] =
    firstPartRegEx.exec(firstPartOfLine) ?? [];

  return {
    lowestAmount: +lowestAmount,
    highestAmount: +highestAmount,
    givenLetter,
    password,
  };
}

function isValidPolicyPartOne(policy: PasswordPolicy): boolean {
  const policyRegEx = new RegExp(`[^${policy.givenLetter}]`, "g");

  const numberOfGivenLetter = policy.password.replace(policyRegEx, "").length;

  return (
    numberOfGivenLetter >= policy.lowestAmount &&
    numberOfGivenLetter <= policy.highestAmount
  );
}

function isValidPolicyPartTwo(policy: PasswordPolicy): boolean {
  const firstLetter = policy.password[policy.lowestAmount - 1];
  const secondLetter = policy.password[policy.highestAmount - 1];

  return (
    firstLetter !== secondLetter &&
    (firstLetter === policy.givenLetter || secondLetter === policy.givenLetter)
  );
}

interface PasswordPolicy {
  lowestAmount: number;
  highestAmount: number;
  givenLetter: string;
  password: string;
}

main();
