package tokens

type Symbol struct {
	Name          Type
	IsNextRequire bool
}

// list of symbols
var Symbols = map[string]Symbol{
	"-": {Name: MINUS, IsNextRequire: false},
	"+": {Name: PLUS, IsNextRequire: false},
	"*": {Name: STAR, IsNextRequire: false},
	"/": {Name: SLASH, IsNextRequire: false},

	"(": {Name: LPARENT, IsNextRequire: false},
	")": {Name: RPARENT, IsNextRequire: false},
	"[": {Name: LBRACKET, IsNextRequire: false},
	"]": {Name: RBRACKET, IsNextRequire: false},

	">":  {Name: GT, IsNextRequire: true},
	">=": {Name: GE, IsNextRequire: false},

	"<":  {Name: LT, IsNextRequire: true},
	"<=": {Name: LE, IsNextRequire: false},
	"<>": {Name: NE, IsNextRequire: false},

	":":  {Name: COLON, IsNextRequire: true},
	":=": {Name: ASSIGN, IsNextRequire: false},

	"=": {Name: EQ, IsNextRequire: false},
	",": {Name: COMMA, IsNextRequire: false},
	";": {Name: SEMICOLON, IsNextRequire: false},
}
