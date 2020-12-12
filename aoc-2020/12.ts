import { importFromFile } from "./util";

function main() {
  const input = importFromFile("12-input.txt");

  const instructions = input.map(mapToInstruction);

  const position = moveToPosition(instructions);

  console.log("Solution Part One: ", calculateManhattan(position));

  const positionWithWaypoint = moveToPositionWithWaypoint(instructions);

  console.log("Solution Part Two: ", calculateManhattan(positionWithWaypoint));
}

function calculateManhattan(position: Position): number {
  return Math.abs(position.eastWest) + Math.abs(position.northSouth);
}

function mapToInstruction(line: string): Instruction {
  const action: any = line[0];
  const amount = +line.slice(1);
  return { action, amount };
}

function moveToPosition(instructions: Instruction[]): Position {
  const position: Position = { northSouth: 0, eastWest: 0 };
  let facing = 0;

  for (const { action, amount } of instructions) {
    switch (action) {
      case "N":
        position.northSouth += amount;
        break;
      case "S":
        position.northSouth -= amount;
        break;
      case "E":
        position.eastWest += amount;
        break;
      case "W":
        position.eastWest -= amount;
        break;
      case "L":
        facing -= amount;
        break;
      case "R":
        facing += amount;
        break;
      case "F":
        switch (facing) {
          case 0:
            position.eastWest += amount;
            break;
          case 90:
            position.northSouth -= amount;
            break;
          case 180:
            position.eastWest -= amount;
            break;
          case 270:
            position.northSouth += amount;
            break;
          default:
            throw new Error("Should never happen");
        }
        break;
      default:
        throw new Error("Should never happen");
    }

    facing = (facing + 360) % 360;
  }

  return position;
}

function moveToPositionWithWaypoint(instructions: Instruction[]): Position {
  const shipPosition: Position = { northSouth: 0, eastWest: 0 };
  let waypointPosition: Position = { northSouth: 1, eastWest: 10 };

  for (const { action, amount } of instructions) {
    switch (action) {
      case "N":
        waypointPosition.northSouth += amount;
        break;
      case "S":
        waypointPosition.northSouth -= amount;
        break;
      case "E":
        waypointPosition.eastWest += amount;
        break;
      case "W":
        waypointPosition.eastWest -= amount;
        break;
      case "F":
        shipPosition.northSouth += waypointPosition.northSouth * amount;
        shipPosition.eastWest += waypointPosition.eastWest * amount;
        break;
      case "R":
      case "L":
        if (
          (amount === 90 && action === "R") ||
          (amount === 270 && action === "L")
        ) {
          waypointPosition = {
            northSouth: -1 * waypointPosition.eastWest,
            eastWest: waypointPosition.northSouth,
          };
        } else if (
          (amount === 90 && action === "L") ||
          (amount === 270 && action === "R")
        ) {
          waypointPosition = {
            northSouth: waypointPosition.eastWest,
            eastWest: -1 * waypointPosition.northSouth,
          };
        } else if (amount === 180) {
          waypointPosition.northSouth *= -1;
          waypointPosition.eastWest *= -1;
        } else {
          throw new Error(
            `Should not happen, but got values: ${action}, ${amount}`
          );
        }

        break;
      default:
        throw new Error("Should never happen");
    }
  }

  return shipPosition;
}

interface Instruction {
  action: "N" | "S" | "E" | "W" | "L" | "R" | "F";
  amount: number;
}

interface Position {
  northSouth: number;
  eastWest: number;
}

main();
