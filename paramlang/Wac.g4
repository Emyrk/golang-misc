// Wac.g4
grammar Wac;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';

fragment DECIMALS
    : [0-9]+
    ;

fragment EXPONENT
    : [eE] [+-]? DECIMALS
    ;

INTEGER: DECIMALS;

// Taken from https://github.com/antlr/grammars-v4/blob/master/golang/GoLexer.g4
FLOAT_LIT              : DECIMALS ('.' DECIMALS? EXPONENT? | EXPONENT)
                       | '.' DECIMALS EXPONENT?
                       ;

fragment OCTAL_DIGIT
    : [0-7]
    ;
fragment HEX_DIGIT
    : [0-9a-fA-F]
    ;

fragment ESCAPED_VALUE
    : '\\' ('u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | 'U' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT
           | [abfnrtv\\'"]
           | OCTAL_DIGIT OCTAL_DIGIT OCTAL_DIGIT
           | 'x' HEX_DIGIT HEX_DIGIT)
    ;
INTERPRETED_STRING_LIT : '"' (~["\\] | ESCAPED_VALUE)*  '"';


BOOLEAN: ('true' | 'false');
WHITESPACE: [ \r\n\t]+ -> skip;
IF: 'if';
ELSE: 'else';
THEN: 'then';
INT_CAST: 'int';
FLOAT_CAST: 'float';

// DELIMITER is the delim for separting hierarchal keys
// Eg: workspace.spec
DELIMITER: '.';
fragment LETTER: [a-zA-Z];
fragment SPECIAL_CHAR: [-_];


// Key_Name is any valid yaml key that we will accept.
// We have to apply some rules
fragment KEY_NAME: LETTER ( LETTER | INTEGER | SPECIAL_CHAR )*;
// After the delimiter, integers are allowed in the keyname to be first
KEY_PATH: KEY_NAME (DELIMITER  (KEY_NAME | INTEGER)+)*;


// Rules

// Key is a full path to a key.
// Eg: workspace.spec
start : statement EOF;

key: KEY_PATH;

literal
   : INTEGER                            # IntegerLit
   | BOOLEAN                            # BoolLit
   | FLOAT_LIT                          # FloatLit
   | INTERPRETED_STRING_LIT             # StringLit
   ;

statement
   : expression                                        # ExpStmt
   | IF expression THEN expression ELSE expression     # IfStmt
   ;

expression
   : literal                            # LiteralExp
   | INT_CAST '(' expression ')'        # IntCast
   | FLOAT_CAST '(' expression ')'      # FloatCast
   | '(' expression ')'                 # Paren
   | expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | expression op=('==' | '!=' | '<' | '<=' | '>' | '>=') expression # Comparator
   | expression op=('and' | 'or') expression # Logical
   | '!' expression                       # InvertLogical
   | key                                  # KeyPath
   ;