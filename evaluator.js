const { SymbolTable, Scope } = require("./symbol_table");
const { ASTNodes } = require("./ast");

const symbolTable = new SymbolTable();

const Evaluator = (ast) => {
    if (!ast) {
        return null;
    }

    if (ast.type === ASTNodes.Literal) {
        return ast.value;
    } else if (ast.type === ASTNodes.Identifier) {
        return symbolTable.lookup(ast.name);
    } else if (ast.type === ASTNodes.Condition) {
        if (Evaluator(ast.condition)) {
            return Evaluator(ast.then);
        } else {
            return Evaluator(ast.el);
        }
    } else if (ast.type === ASTNodes.Abstraction) {
        const scope = new Scope();
		return x => {
			scope.add(ast.arg.id.name, x);
			symbolTable.push(scope);
			return Evaluator(ast.body);
		};
    } else if (ast.type === ASTNodes.IsZero) {
        return Evaluator(ast.expression) === 0;
    } else if (ast.type === ASTNodes.Arithmetic) {
        const op = ast.operator;
        const val = Evaluator(ast.expression);
        switch (op) {
            case 'succ':
                return val + 1;
            case 'pred':
                return(val - 1 >= 0) ? val + 1 : val;    
        }
    } else if (ast.type === ASTNodes.Application) {
        const left = Evaluator(ast.left);
        const right = Evaluator(ast.right);
        return left(right);
    }
    return true;
}

module.exports.Evaluator = Evaluator;
