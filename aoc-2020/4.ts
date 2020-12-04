import { importFromFile } from "./util";

function main() {
  const input = importFromFile("4-input.txt");

  const passports = mapToPassports(input);

  const validPassportsPartOne = passports.filter(isValidPassportPartOne);

  console.log("Solution Part One: ", validPassportsPartOne.length);

  const validPassportPartTwo = passports.filter(isValidPassportPartTwo);

  console.log("Solution Part Two: ", validPassportPartTwo.length);
}

function mapToPassports(input: string[]): Passport[] {
  const passports: Passport[] = [];

  let tempString = "";

  for (const line of input) {
    if (!!line) {
      tempString += line + " ";
    } else {
      tempString.trim();
      const passport = mapToPassport(tempString);
      passports.push(passport);
      tempString = "";
    }
  }
  return passports;
}

function mapToPassport(str: string): Passport {
  const parts = str.split(" ").filter((s) => !!s);
  const passportArr = parts.map((part) => part.split(":"));

  return Object.fromEntries(passportArr);
}

function isValidPassportPartOne(passport: Passport): boolean {
  const requiredFields = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"];

  return requiredFields.every((field) => !!passport[field]);
}

function isValidPassportPartTwo(passport: Passport): boolean {
  if (!isValidPassportPartOne(passport)) {
    return false;
  }

  const birthYear = +passport.byr;
  if (birthYear < 1920 || birthYear > 2002) {
    return false;
  }

  const issueYear = +passport.iyr;
  if (issueYear < 2010 || issueYear > 2020) {
    return false;
  }

  const expirationYear = +passport.eyr;
  if (expirationYear < 2020 || expirationYear > 2030) {
    return false;
  }

  const heightRegEx = /(\d+)(in|cm)?/;
  const [, height, unit] = heightRegEx.exec(passport.hgt) ?? [];
  if (
    !unit ||
    (unit === "cm" && +height < 150) ||
    (unit === "cm" && +height > 193) ||
    (unit === "in" && +height < 59) ||
    (unit === "in" && +height > 76)
  ) {
    return false;
  }

  const hairColorRegex = /#[0-9a-f]{6}/;
  if (!hairColorRegex.test(passport.hcl)) {
    return false;
  }

  const eyeColors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];
  if (!eyeColors.some((color) => passport.ecl === color)) {
    return false;
  }

  const passportIDRegex = /\d{9}/;
  if (!passportIDRegex.test(passport.pid) || passport.pid.length > 9) {
    return false;
  }

  return true;
}

interface Passport {
  [index: string]: string | undefined;
  byr: string;
  iyr: string;
  eyr: string;
  hgt: string;
  hcl: string;
  ecl: string;
  pid: string;
  cid?: string;
}

main();
