import { importFromFile } from "./util";

function main() {
  const input = importFromFile("7-input.txt");
  const bagRules = input.reduce(mapToBagRules, new Map());

  const containingBags = findContainingBags("shiny gold", bagRules);

  // Shiny gold bags can't contain themselves
  console.log("Solution Part One: ", containingBags.length - 1);

  const numberOfBagsContainedInShinyGold = findAmountOfBags(
    "shiny gold",
    bagRules
  );

  console.log("Solution Part Two: ", numberOfBagsContainedInShinyGold);
}

function mapToBagRules(prev: BagRules, current: string): BagRules {
  const bagRegEx = /(.+) bags contain(?: (\d.+)? bags?(?:,|\.)?)+/;
  if (current.includes("no other bags")) {
    return prev;
  }
  const [, keyColor, containedBagsStr] = bagRegEx.exec(current) ?? [];
  const containedBags = mapToContainedBags(containedBagsStr);

  prev.set(keyColor, containedBags);

  return prev;
}

function mapToContainedBags(str: string): BagRule[] {
  const parts = str.replace(/ bags?/g, "").split(", ");
  const rules = parts.map((part) => {
    const amount = +part[0];
    const color = part.slice(2);
    return {
      amount,
      color,
    };
  });
  return rules;
}

function findContainingBags(
  colorToSearchFor: string,
  bagRules: BagRules
): string[] {
  const foundColors = [];
  for (const color of bagRules.keys()) {
    if (containsBag(color, colorToSearchFor, bagRules)) {
      foundColors.push(color);
    }
  }

  return foundColors;
}

function containsBag(
  color: string,
  colorToSearchFor: string,
  bagRules: BagRules
): boolean | undefined {
  if (!bagRules.has(color)) {
    return;
  }

  if (color === colorToSearchFor) {
    return true;
  }

  for (const rule of bagRules.get(color) ?? []) {
    const isContained = containsBag(rule.color, colorToSearchFor, bagRules);

    if (isContained) {
      return true;
    }
  }

  return;
}

function findAmountOfBags(color: string, bagRules: BagRules): number {
  let amount = 0;
  if (!bagRules.has(color)) {
    return 0;
  }

  for (const rule of bagRules.get(color) ?? []) {
    amount +=
      rule.amount + rule.amount * findAmountOfBags(rule.color, bagRules);
  }

  return amount;
}

type BagRules = Map<string, BagRule[]>;

interface BagRule {
  color: string;
  amount: number;
}

main();
