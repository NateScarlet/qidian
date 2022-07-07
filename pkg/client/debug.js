const fs = require("fs");

let evalCount = 0;
const readFromFile = false;

function onEval(code) {
  evalCount += 1;
  const filename = `eval_${evalCount}.local.js`;
  if (readFromFile && fs.existsSync(filename)) {
    return fs.readFileSync(filename);
  }
  fs.writeFileSync(filename, code);
  return code;
}

module.exports = {
  onEval,
};
