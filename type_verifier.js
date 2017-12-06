const { SymbolTable, Scope } = require('./symbol_table');
const { ASTNodes, Types } = require('./ast');

const symbolTable = new SymbolTable();

const TypeEq = (a, b) => {
    if (a instanceof Array && b instanceof Array) {
        if (a.length != b.length) {
            return false;
        } else {
            for (let i = 0; i < a.length; i++) {
                if (!TypeEq(a[i], b[i])) {
                    return false;
                }
            }
            return true;
        }
    } else {
        if (typeof a === 'string' && typeof b === 'string') {
            return a === b;
        }
    }
    return false;
}

const TypeCheck = (ast, diagnostics) => {
    diagnostics = diagnostics || [];

    if (!ast) {
        return true;
    }

    if (ast.type === ASTNodes.Literal) {
        if (ast.value == 0) {
            return {
                diagnostics,
                type: Types.Natural
            };
        } else if (ast.value == false || ast.value == true) {
            return {
                diagnostics,
                type: Types.Boolean
            };
        } else {
            diagnostics.push('Unknown literal value');
        }
    }
    
    if (ast.type === ASTNodes.Identifier) {
        return {
            diagnostics,
            type: symbolTable.lookup(ast.name)
        };
    }

    if (ast.type === ASTNodes.Condition) {
        if (!ast.then || !ast.el || !ast.condition) {
            diagnostics.push('If statement ill formed');
            return { diagnostics };
        }
        const c = TypeCheck(ast.condition);
        diagnostics = diagnostics.concat(c.diagnostics);
        if (!TypeEq(c.type, ASTNodes.Boolean)) {
            diagnostics.push('If condition has wrong type');
            return { diagnostics };
            
        }
        const t = TypeCheck(ast.then);
        diagnostics = diagnostics.concat(t.diagnostics);
        const e = TypeCheck(ast.else);
        diagnostics = diagnostics.concat(e.diagnostics);
        if (TypeEq(t.type, e.type)) {
            return t.type;
        } else {
            diagnostics.push('Type mismatch of then and else expression');
            return { diagnostics };
        }
    }

    if (ast.type === ASTNodes.Abstraction) {
        const scope = new Scope();
        scope.add(ast.arg.id.name, ast.arg.type);
		symbolTable.push(scope);
		if (!ast.body) {
			diagnostics.push("No body of a function");
			return { diagnostics };
        }
        
        const body = TypeCheck(ast.body);
		const bodyType = body.type;
		diagnostics = diagnostics.concat(body.diagnostics);
		if (!bodyType) {
			diagnostics.push("Incorrect type of the body");
			return { diagnostics };
		}
        return { diagnostics, type: [ast.arg.type, bodyType] };
    }

    if (ast.type === ASTNodes.IsZero) {
        const body = TypeCheck(ast.expression);
        diagnostics = diagnostics.concat(body.diagnostics);
        if (!TypeEq(body.type, Types.Natural)) {
            diagnostics.push('Incorrect type for iszero function argument');
            return { diagnostics };
        }
        return { diagnostics, type: Types.Boolean };
    }

    if (ast.type === ASTNodes.Arithmetic) {
        const body = TypeCheck(ast.expression);
		diagnostics = diagnostics.concat(body.diagnostics);
		const bodyType = body.type;
		if (!TypeEq(bodyType, Types.Natural)) {
			diagnostics.push(`Incorrect type of ${ast.operation}`);
			return { diagnostics };
		}
		return { diagnostics, type: Types.Natural };
    }

    if (ast.type === ASTNodes.Application) {
        const left = TypeCheck(ast.left);
        const leftType = left.type || [];
        diagnostics = diagnostics.concat(left.diagnostics);
        const right = TypeCheck(ast.right);
        const rightType = right.type || [];
        diagnostics = diagnostics.concat(right.diagnostics);

        if (leftType.length) {
            if (!ast.right || leftType[0] == rightType) {
                return {
                    diagnostics,
                    type: leftType[1]
                };
            } else {
                diagnostics.push('Incorrect application type');
                return { diagnostics };
            }
        } else {
            return { diagnostics };
        }
    }
    return { diagnostics };
}

module.exports.TypeCheck = ast => TypeCheck(ast);



