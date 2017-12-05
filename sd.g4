grammar sd;

expr            : ID                                    #symbol
                | DIGIT                                 #number
                | 'λ' ID ':' r_type ('=>'|'→') expr     #abstraction
                | expr expr                             #applicatio
                | '(' expr ')'                          #expression
                | arithmeticOps expr                    #arithmeicExps
                | ISZERO expr                           #checkZero
                | ifThen                                #ifElseExpr
                ;

arithmeticOps   : SUCC
                | PRED
                ;

ifThen          : IF expr THEN expr ELSE expr
                ;

r_type          : NAT
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