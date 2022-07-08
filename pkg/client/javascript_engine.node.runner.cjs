// @ts-check
/// <reference no-default-lib="true" />
/// <reference lib="ES2016" />
/// <reference types="./javascript_engine.node.env"/>

const fs = require("fs");
const vm = require("vm");

let evalCount = 0;

function onEval(code) {
  const readFromFile = false;
  if (__DEBUG__) {
    evalCount += 1;
    const filename = `eval_${evalCount}.local.js`;
    if (evalCount > 1 && readFromFile && fs.existsSync(filename)) {
      return fs.readFileSync(filename);
    }
    fs.writeFileSync(filename, code);
  }
  return code;
}

const ctx = vm.createContext({ eval: vmEval, escape, unescape });

function vmEval(code) {
  code = onEval(code);
  return vm.runInContext(code, ctx);
}

(async () => {
  console.log(await vmEval(__CODE__));
})().catch((err) => {
  console.error(err);
  process.exit(1);
});
