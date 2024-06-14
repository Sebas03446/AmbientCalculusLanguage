grammar AmbientCalculus;

program: statement* EOF;

statement: processDeclaration
         | ambientDeclaration
         | movementStatement
         | communicationStatement
         | inStatement
         | outStatement
         | openStatement
         | callStatement
         | printStatement
         | ifStatement
         | variableDeclaration
         ;

processDeclaration: 'process' ID '{' processStatement* '}';

processStatement: statement
                | moveStatement;

moveStatement: 'move' AMBIENTID ';';

ambientDeclaration: 'ambient' AMBIENTID '{' conditions? statement* '}';

conditions: 'conditions' '{' statement* '}';

movementStatement: 'move' ID 'to' AMBIENTID ';';

communicationStatement: 'send' expression 'to' ID ';'
                      | 'receive' ID 'from' ID ';'
                      ;

inStatement: 'in' AMBIENTID ';';
outStatement: 'out' ';';
openStatement: 'open' AMBIENTID ';';
callStatement: 'call' ID ';';
printStatement: 'print' expression ';';
ifStatement: 'if' '(' expression ')' '{' statement* '}';

variableDeclaration: 'let' ID '=' expression ';';

expression: expression operator expression
          | '(' expression ')'
          | ID
          | INT
          | STRING
          ;

operator: '+' | '-' | '*' | '/';

AMBIENTID: [a-z][a-z0-9_]*;
ID: [A-Z][A-Z0-9_]*;

INT: [0-9]+;

STRING: '"' (~["\r\n])* '"';

WS: [ \t\r\n]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
