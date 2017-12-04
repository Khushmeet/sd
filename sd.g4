grammar sd;

program         : application
                ;

application     : exprAbs application+
                ;

exprAbs         : expr
                | abstraction
                ;


abstraction     : '(' 'λ' ID ':' type '=>' application ')'
                ;

expr            : ifThen
                | isZeroCheck
                | arithmeticOps
                | ZERO
                | TRUE
                | FALSE
                | ID
	            | paramExpr
                ;

// TODO: Check for reserved words
ID              : [a-z]+
                ;

type            : NAT
                | BOOL
                ;

ifThen          : IF application THEN application ELSE application
                ;

isZeroCheck     : ISZERO application
                ;

arithmeticOps   : operator application
                ;

paramExpr       : '(' expr ')'
                ;

operator        : SUCC
                | PRED
                ;

RESERVEDWORDS   : IF
                | THEN
                | ELSE
                | PRED
                | SUCC
                | TYPES
                | ISZERO
                | TRUE
                | FALSE
                ;

TYPES           : NAT
                | BOOL
                ;


ZERO            : '0'
                ;

SUCC            : 'succ'
                ;

PRED            : 'pred'
                ;

TRUE            : 'true'
                ;

FALSE           : 'false'
                ;

IF              : 'if'
                ;

THEN            : 'then'
                ;

ELSE            : 'else'
                ;

ISZERO          : 'iszero'
                ;

NAT             : 'Nat'
                ;

BOOL            : 'Bool'
                ;

WS              : [ \t\r\n]+ -> skip
                ;