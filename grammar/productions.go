package grammar

type Production struct {
	head string
	body string
}

// !!!!!!!!!!!!!!!!!!!!!!!! Don't change these rules !!!!!!!!!!!!!!!!!!!!!!!!!!!
var Productions = map[string]Production{
	"0": {
		head: "P'",
		body: "P",
	},
	"1": {
		head: "P",
		body: "program id output ; declarations subprogram_declaration begin optional_statment end.",
	},
	"2": {
		head: "output",
		body: "( id_list )",
	},
	"3": {
		head: "P",
		body: "program id . output ; declarations subprogram_declaration begin optional_statment end.",
	},
	"4": {
		head: "id_list",
		body: "id",
	},
	"5": {
		head: "id_list",
		body: "id_list , id",
	},
	"6": {
		head: "declarations",
		body: "declarations var id_list : type ;",
	},
	"7": {
		head: "declarations",
		body: "",
	},
	"8": {
		head: "type",
		body: "standard_type",
	},
	"9": {
		head: "type",
		body: "array [ int .. int ] of standard_type",
	},
	"10": {
		head: "standard_type",
		body: "integer",
	},
	"11": {
		head: "standard_type",
		body: "real",
	},
	"12": {
		head: "standard_type",
		body: "string",
	},
	"13": {
		head: "standard_type",
		body: "boolean",
	},
	"14": {
		head: "subprogram_declaration",
		body: "subprogram_declaration -> subprogram_declarations subprogram_declaration ;",
	},
	"15": {
		head: "subprogram_declaration",
		body: "",
	},
	"16": {
		head: "subprogram_declarations",
		body: "supprogram_head declarations compound_statment",
	},
	"17": {
		head: "supprogram_head",
		body: "function id arguments : standard_type ;",
	},
	"18": {
		head: "supprogram_head",
		body: "procedure id arguments ;",
	},
	"19": {
		head: "arguments",
		body: "( parameter_list )",
	},
	"20": {
		head: "arguments",
		body: "",
	},
	"21": {
		head: "parameter_list",
		body: "id_list : type",
	},
	"22": {
		head: "parameter_list",
		body: "parameter_list ; id_list : type",
	},
	"23": {
		head: "compound_statment",
		body: "begin optional_statment end",
	},
	"24": {
		head: "optional_statment",
		body: "statment_list",
	},
	"25": {
		head: "optional_statment",
		body: "",
	},
	"26": {
		head: "statment_list",
		body: "statment",
	},
	"27": {
		head: "statment_list",
		body: "statment_list ; optional_statment",
	},
	"28": {
		head: "statment",
		body: "varible := expr",
	},
	"29": {
		head: "statment",
		body: "procedure_statment",
	},
	"30": {
		head: "statment",
		body: "compound_statment",
	},
	"31": {
		head: "statment",
		body: "if expr then statment",
	},
	"32": {
		head: "statment",
		body: "if expr then statment else statment",
	},
	"33": {
		head: "statment",
		body: "while expr do statment",
	},
	"34": {
		head: "statment",
		body: "for varible := expr st_for st_for2 do statment",
	},
	"35": {
		head: "st_for",
		body: "to",
	},
	"36": {
		head: "st_for",
		body: "downto",
	},
	"37": {
		head: "st_for2",
		body: "id",
	},
	"38": {
		head: "st_for2",
		body: "int",
	},
	"39": {
		head: "varible",
		body: "id",
	},
	"40": {
		head: "varible",
		body: "id [ expr ]",
	},
	"41": {
		head: "procedure_statment",
		body: "id",
	},
	"42": {
		head: "procedure_statment",
		body: "id ( expr_list )",
	},
	"43": {
		head: "expr_list",
		body: "expr",
	},
	"44": {
		head: "expr_list",
		body: "expr_list , expr",
	},
	"45": {
		head: "expr_list",
		body: "",
	},
	"46": {
		head: "expr",
		body: "simple_expr",
	},
	"47": {
		head: "expr",
		body: "simple_expr relop simple_expr",
	},
	"48": {
		head: "relop",
		body: "<",
	},
	"49": {
		head: "relop",
		body: "<=",
	},
	"50": {
		head: "relop",
		body: "<>",
	},
	"51": {
		head: "relop",
		body: "=",
	},
	"52": {
		head: "relop",
		body: ">=",
	},
	"53": {
		head: "relop",
		body: ">",
	},
	"54": {
		head: "simple_expr",
		body: "term",
	},
	"55": {
		head: "simple_expr",
		body: "sign term",
	},
	"56": {
		head: "simple_expr",
		body: "simple_expr addop term",
	},
	"57": {
		head: "addop",
		body: "+",
	},
	"58": {
		head: "addop",
		body: "-",
	},
	"59": {
		head: "addop",
		body: "or",
	},
	"60": {
		head: "term",
		body: "factor",
	},
	"61": {
		head: "term",
		body: "term mulop factor",
	},
	"62": {
		head: "mulop",
		body: "*",
	},
	"63": {
		head: "mulop",
		body: "/",
	},
	"64": {
		head: "mulop",
		body: "div",
	},
	"65": {
		head: "mulop",
		body: "mod",
	},
	"66": {
		head: "mulop",
		body: "and",
	},
	"67": {
		head: "factor",
		body: "id",
	},
	"68": {
		head: "factor",
		body: "id ( expr_list )",
	},
	"69": {
		head: "factor",
		body: "int",
	},
	"70": {
		head: "factor",
		body: "float",
	},
	"71": {
		head: "factor",
		body: "string",
	},
	"72": {
		head: "factor",
		body: "boolean",
	},
	"73": {
		head: "factor",
		body: "( expr )",
	},
	"74": {
		head: "factor",
		body: "not factor",
	},
	"75": {
		head: "sign",
		body: "+",
	},
	"76": {
		head: "sign",
		body: "-",
	},
}
