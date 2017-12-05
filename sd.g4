grammar sd;

expr            : ID
                | DIGIT
                | 'λ' ID ':' type ('=>'|'→') expr
                | expr expr
                | '(' expr ')'
                | arithmeticOps expr
                | ISZERO expr
                | ifThen
                ;

arithmeticOps   : SUCC
                | PRED
                ;

ifThen          : IF expr THEN expr ELSE expr
                ;

type            : NAT
                | BOOL
                ;

IF              : 'if'
                ;

ELSE            : 'else'
                ;

THEN            : 'then'
                ;

ISZERO          : 'iszero'
                ;

SUCC            : 'succ'
                ;

PRED            : 'pred'
                ;

ID              : [a-z]+
                ;

DIGIT           : [0-9]+
                ;

NAT             : 'Nat'
                ;

BOOL            : 'Bool'
                ;

WS              : [ \t\r\n]+    -> skip
                ;