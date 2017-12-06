grammar sd;


expr            : TRUE	                                #trueLiteral
	            | FALSE	                                #falseLiteral 
                | ID                                    #symbol
                | DIGIT                                 #number
                | 'λ' ID ('=>'|'→') expr                #abstraction
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

TRUE            : 'true'
                ;

FALSE           : 'false'
                ;

WS              : [ \t\r\n]+    -> skip
                ;