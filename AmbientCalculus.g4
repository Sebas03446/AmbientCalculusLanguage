grammar AmbientCalculus;

program: statement* EOF;

statement: processDeclaration
         | ambientDeclaration
         | movementStatement
         | sendStatement
         | receiveStatement
         | inStatement
         | outStatement
         | openStatement
         | callStatement
         | printStatement
         | variableDeclaration
         | assignmentStatement
         | modifyConditionStatement
         ;

processDeclaration: 'process' ID '{' processStatement* '}';

processStatement: statement
                | moveStatement
                | printConditionStatement;

moveStatement: 'move' AMBIENTID ';';

ambientDeclaration: 'ambient' AMBIENTID '{' conditions? statement* '}';

conditions: conditionsBlock '{' conditionsVariableDeclaration*  '}';

movementStatement: 'move' ID 'to' AMBIENTID ';';

sendStatement: 'send' STRING 'to' ID ';';
receiveStatement: 'receive'';';

inStatement: 'in' AMBIENTID ';';
outStatement: 'out' ';';
openStatement: 'open' AMBIENTID ';';
callStatement: 'call' ID ';';
printStatement: 'print' expression ';';
printConditionStatement: 'printc' ID ';';

modifyConditionStatement: 'modifyc' ID op=('+=' | '-=') expression ';';
conditionsBlock: 'conditions';
variableDeclaration: 'let' ID '=' expression ';';

conditionsVariableDeclaration: 'letc' ID '=' INT restriction? ';';
restriction: 'restriction' comparator INT;

assignmentStatement: ID '=' expression ';';

expression: expression operator expression
          | '(' expression ')'
          | ID
          | INT
          | STRING
          ;

operator: '+' | '-' | '*' | '/';

comparator: '>' | '<' | '>=' | '<=' | '==' | '!=';


AMBIENTID: [a-z][a-z0-9_]*;
ID: [A-Z][A-Z0-9_]*;

INT: [0-9]+;

STRING: '"' (~["\r\n])* '"';

WS: [ \t\r\n]+ -> skip;
COMMENT: '//' ~[\r\n]* -> skip;
