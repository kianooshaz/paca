package tokens

// list of keywords
var Keywords = map[string]Type{
	"do":        DO,
	"while":     WHILE,
	"else":      ELSE,
	"then":      THEN,
	"if":        IF,
	"end":       END,
	"begin":     BEGIN,
	"function":  FUNCTION,
	"procedure": PROCEDURE,
	"real":      REAL,
	"integer":   INTEGER,
	"and":       AND,
	"mod":       MOD,
	"or":        OR,
	"not":       NOT,
	"of":        OF,
	"array":     ARRAY,
	"var":       VAR,
	"program":   PROGRAM,
	"downto":    DOWNTO,
	"to":        TO,
	"for":       FOR,
	"id":        ID,
	"int":       INT,
	"float":     FLOAT,
	"string":    STRING,
	"boolean":   BOOLEAN,
	"div":       DIV,
}
