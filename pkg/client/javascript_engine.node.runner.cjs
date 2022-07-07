/// <reference no-default-lib="true" />
/// <reference types="./javascript_engine.node.env"/>

const vm = require("vm");
const ctx = vm.createContext({ eval: vmEval, escape, unescape });
function vmEval(code) {
  if (__DEBUG__) {
    const debug = require("./debug");
    code = debug.onEval(code);
  }
  return vm.runInContext(code, ctx);
}
ctx.eval = vmEval;
(async () => {
  console.log(await vmEval(__CODE__));
})().catch((err) => {
  console.error(err);
  process.exit(1);
});
