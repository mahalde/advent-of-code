import { readFileSync } from "fs";

export function importFromFile(filepath: string): string[] {
  const fileBuffer = readFileSync(filepath);
  const file = fileBuffer.toString("utf-8");
  return file.split("\r\n");
}
