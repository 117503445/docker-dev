#!/usr/bin/env node

const fs = require("fs");
const os = require("os");
const path = require("path");

const tool = process.argv[2] || "unknown";
const cwd = process.env.CLAUDE_PROJECT_DIR || process.cwd();
const logDir = path.join(os.homedir(), ".ai", encodeURIComponent(cwd));
const logPath = path.join(logDir, `${tool}.jsonl`);

fs.mkdirSync(logDir, { recursive: true });

let raw = "";
process.stdin.setEncoding("utf8");
process.stdin.on("data", (chunk) => {
  raw += chunk;
});
process.stdin.on("end", () => {
  fs.appendFileSync(logPath, raw);
  if (!raw.endsWith("\n")) {
    fs.appendFileSync(logPath, "\n");
  }
});
