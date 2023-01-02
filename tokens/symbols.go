package tokens

type Symbol struct {
	Name    Type
	HasNext bool
}

// list of symbols
var Symbols = map[string]Symbol{
	"-": {Name: MINUS, HasNext: false},
	"+": {Name: PLUS, HasNext: false},
	"*": {Name: STAR, HasNext: false},
	"/": {Name: SLASH, HasNext: false},

	"(": {Name: LPARENT, HasNext: true},
	")": {Name: RPARENT, HasNext: false},
	"[": {Name: LBRACKET, HasNext: false},
	"]": {Name: RBRACKET, HasNext: false},

	">":  {Name: GT, HasNext: true},
	">=": {Name: GE, HasNext: false},

	"<":  {Name: LT, HasNext: true},
	"<=": {Name: LE, HasNext: false},
	"<>": {Name: NE, HasNext: false},

	":":  {Name: COLON, HasNext: true},
	":=": {Name: ASSIGN, HasNext: false},

	"=": {Name: EQ, HasNext: false},
	",": {Name: COMMA, HasNext: false},
	";": {Name: SEMICOLON, HasNext: false},
	".": {Name: DOT, HasNext: true},
}
