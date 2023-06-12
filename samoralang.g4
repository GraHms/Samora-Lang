grammar samoralang;

program           : statement* ;

statement         : variableDeclaration
                  | assignmentStatement
                  | printStatement
                  | readStatement
                  | ifStatement
                  | whileStatement
                  | functionDeclaration ;

variableDeclaration : 'let' Identifier '=' expression ;

assignmentStatement : Identifier '=' expression ;

printStatement    : 'print' '(' expression ')' ;
readStatement     : 'read' '(' Identifier ')' ;

ifStatement       : 'if' '(' expression ')' '{' statement* '}' ('else' '{' statement* '}')? ;

whileStatement    : 'while' '(' expression ')' '{' statement* '}' ;

functionDeclaration : 'let' Identifier '=' 'fn' '(' parameterList? ')' '{' statement* '}' ;

parameterList     : Identifier (',' Identifier)* ;

expression        : logicalOrExpression ;

logicalOrExpression : logicalAndExpression ('||' logicalAndExpression)* ;
logicalAndExpression : equalityExpression ('&&' equalityExpression)* ;
equalityExpression : relationalExpression ( ('==' | '!=') relationalExpression)* ;
relationalExpression : additiveExpression (('<' | '>' | '<=' | '>=') additiveExpression)* ;
additiveExpression : multiplicativeExpression (('+' | '-') multiplicativeExpression)* ;
multiplicativeExpression : unaryExpression (('*' | '/') unaryExpression)* ;
unaryExpression   : ('+' | '-' | '!') unaryExpression
                  | primaryExpression ;

primaryExpression : Literal
                  | Identifier
                  | functionCall
                  | '(' expression ')' ;

functionCall      : Identifier '(' argumentList? ')' ;

argumentList      : expression (',' expression)* ;

Identifier        : [a-zA-Z_][a-zA-Z0-9_]* ;
Literal           : StringLiteral | NumericLiteral ;

StringLiteral     : '"' .*? '"' ;
NumericLiteral    : [0-9]+ ;
