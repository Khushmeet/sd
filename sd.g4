grammar sd;

program         : '('? application ')'?
                ;

application     : exprAbs application
                ;

exprAbs         : expr
                | abstraction
                ;


abstrcation     : '(' 'λ' id ':' type ('→'|'=>') application ')'
                ;

expr            : ifThen
                | isZeroCheck
                | arithmeticOps
                | id
                | paramExpr
                | ZERO
                | TRUE
                | FALSE
                ;

// TODO: Check for reserved words
id              : [a-z]+
                ;

type            : 'Nat'
                | 'Bool'
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

ZERO            : [0]
                ;

Succ            : 'succ'
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

ISZERO          : 'isZero'
                ;
