package tokens

const (
	EOF        Type = "$"
	MINUS      Type = "-"
	NOT        Type = "not"
	PLUS       Type = "+"
	RPARENT    Type = ")"
	LPARENT    Type = "("
	AND        Type = "and"
	MOD        Type = "mod"
	DIV        Type = "div"
	SLASH      Type = "/"
	STAR       Type = "*"
	OR         Type = "or"
	GT         Type = ">"
	GE         Type = ">="
	EQ         Type = "="
	NE         Type = "<>"
	LE         Type = "<="
	LT         Type = "<"
	COMMA      Type = ","
	RBRACKET   Type = "]"
	LBRACKET   Type = "["
	DO         Type = "do"
	WHILE      Type = "while"
	ELSE       Type = "else"
	THEN       Type = "then"
	IF         Type = "if"
	ASSIGN     Type = ":="
	SEMICOLON  Type = ";"
	END        Type = "end"
	BEGIN      Type = "begin"
	COLON      Type = ":"
	FUNCTION   Type = "function"
	PROCEDURE  Type = "procedure"
	REAL       Type = "real"
	INTEGER    Type = "integer" // standard type
	DOTDOT     Type = ".."
	OF         Type = "of"
	ARRAY      Type = "array"
	VAR        Type = "var"
	DOT        Type = "."
	PROGRAM    Type = "program"
	ENDSTOP    Type = "end."
	DOWNTO     Type = "downto"
	TO         Type = "to"
	FOR        Type = "for"
	ID         Type = "id"
	INT        Type = "int"
	FLOAT      Type = "float"
	SCIENTIFIC Type = "Scientific"
	STRING     Type = "string"
	BOOLEAN    Type = "boolean"
	IDENT      Type = "id"
	DOLLAR     Type = "$"
)
