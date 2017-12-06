const { parse } = require("./sd");
const { TypeCheck } = require("./type_verifier");
const { Evaluator } = require("./evaluator");
const { green, red } = require("chalk");
const { readFileSync, existsSync } = require("fs");

const getAst = program => {
	const ast = parse(program);

	const diagnostics = TypeCheck(ast).diagnostics;
	if (diagnostics.length) {
		console.error(red(diagnostics.join("\n")));
		process.exit(1);
	}

	return ast;
};

const fileName = process.argv.pop();
let mode = process.argv.pop();
let target = process.argv.pop();

if (!existsSync(fileName)) {
	console.error(`"${fileName}" does not exist.`);
	process.exit(1);
}

const userMessage = `Evaluating ${fileName}`;

console.log(green(userMessage));
console.log();

const program = readFileSync(fileName, { encoding: "utf-8" });
const ast = getAst(program);
console.log(Evaluator(ast));
