package grammar

import (
	"fmt"
	"strings"
)

type Production struct {
	Head string
	Body string
}

func (p *Production) ToString() string {
	body := p.Body
	if body == "" {
		body = "ε"
	}
	return fmt.Sprintf("%s → %s", p.Head, body)
}

func (p *Production) BodySize() int {
	if p.Body == "" {
		return 0
	}
	body := strings.Trim(p.Body, " ")
	return len(strings.Split(body, " "))
}

// !!!!!!!!!!!!!!!!!!!!!!!! Don't change these rules !!!!!!!!!!!!!!!!!!!!!!!!!!!
var Productions = map[string]Production{
	"0": {
		Head: "P'",
		Body: "P",
	},
	"1": {
		Head: "P",
		Body: "program id output ; declarations subprogram_declaration begin optional_statment end.",
	},
	"2": {
		Head: "output",
		Body: "( id_list )",
	},
	"3": {
		Head: "output",
		Body: "",
	},
	"4": {
		Head: "id_list",
		Body: "id",
	},
	"5": {
		Head: "id_list",
		Body: "id_list , id",
	},
	"6": {
		Head: "declarations",
		Body: "declarations var id_list : type ;",
	},
	"7": {
		Head: "declarations",
		Body: "",
	},
	"8": {
		Head: "type",
		Body: "standard_type",
	},
	"9": {
		Head: "type",
		Body: "array [ int .. int ] of standard_type",
	},
	"10": {
		Head: "standard_type",
		Body: "integer",
	},
	"11": {
		Head: "standard_type",
		Body: "real",
	},
	"12": {
		Head: "standard_type",
		Body: "string",
	},
	"13": {
		Head: "standard_type",
		Body: "boolean",
	},
	"14": {
		Head: "subprogram_declaration",
		Body: "subprogram_declarations subprogram_declaration ;",
	},
	"15": {
		Head: "subprogram_declaration",
		Body: "",
	},
	"16": {
		Head: "subprogram_declarations",
		Body: "supprogram_head declarations compound_statment",
	},
	"17": {
		Head: "supprogram_head",
		Body: "function id arguments : standard_type ;",
	},
	"18": {
		Head: "supprogram_head",
		Body: "procedure id arguments ;",
	},
	"19": {
		Head: "arguments",
		Body: "( parameter_list )",
	},
	"20": {
		Head: "arguments",
		Body: "",
	},
	"21": {
		Head: "parameter_list",
		Body: "id_list : type",
	},
	"22": {
		Head: "parameter_list",
		Body: "parameter_list ; id_list : type",
	},
	"23": {
		Head: "compound_statment",
		Body: "begin optional_statment end",
	},
	"24": {
		Head: "optional_statment",
		Body: "statment_list",
	},
	"25": {
		Head: "optional_statment",
		Body: "",
	},
	"26": {
		Head: "statment_list",
		Body: "statment",
	},
	"27": {
		Head: "statment_list",
		Body: "statment_list ; optional_statment",
	},
	"28": {
		Head: "statment",
		Body: "varible := expr",
	},
	"29": {
		Head: "statment",
		Body: "procedure_statment",
	},
	"30": {
		Head: "statment",
		Body: "compound_statment",
	},
	"31": {
		Head: "statment",
		Body: "if expr then statment",
	},
	"32": {
		Head: "statment",
		Body: "if expr then statment else statment",
	},
	"33": {
		Head: "statment",
		Body: "while expr do statment",
	},
	"34": {
		Head: "statment",
		Body: "for varible := expr st_for st_for2 do statment",
	},
	"35": {
		Head: "st_for",
		Body: "to",
	},
	"36": {
		Head: "st_for",
		Body: "downto",
	},
	"37": {
		Head: "st_for2",
		Body: "id",
	},
	"38": {
		Head: "st_for2",
		Body: "int",
	},
	"39": {
		Head: "varible",
		Body: "id",
	},
	"40": {
		Head: "varible",
		Body: "id [ expr ]",
	},
	"41": {
		Head: "procedure_statment",
		Body: "id",
	},
	"42": {
		Head: "procedure_statment",
		Body: "id ( expr_list )",
	},
	"43": {
		Head: "expr_list",
		Body: "expr",
	},
	"44": {
		Head: "expr_list",
		Body: "expr_list , expr",
	},
	"45": {
		Head: "expr_list",
		Body: "",
	},
	"46": {
		Head: "expr",
		Body: "simple_expr",
	},
	"47": {
		Head: "expr",
		Body: "simple_expr relop simple_expr",
	},
	"48": {
		Head: "relop",
		Body: "<",
	},
	"49": {
		Head: "relop",
		Body: "<=",
	},
	"50": {
		Head: "relop",
		Body: "<>",
	},
	"51": {
		Head: "relop",
		Body: "=",
	},
	"52": {
		Head: "relop",
		Body: ">=",
	},
	"53": {
		Head: "relop",
		Body: ">",
	},
	"54": {
		Head: "simple_expr",
		Body: "term",
	},
	"55": {
		Head: "simple_expr",
		Body: "sign term",
	},
	"56": {
		Head: "simple_expr",
		Body: "simple_expr addop term",
	},
	"57": {
		Head: "addop",
		Body: "+",
	},
	"58": {
		Head: "addop",
		Body: "-",
	},
	"59": {
		Head: "addop",
		Body: "or",
	},
	"60": {
		Head: "term",
		Body: "factor",
	},
	"61": {
		Head: "term",
		Body: "term mulop factor",
	},
	"62": {
		Head: "mulop",
		Body: "*",
	},
	"63": {
		Head: "mulop",
		Body: "/",
	},
	"64": {
		Head: "mulop",
		Body: "div",
	},
	"65": {
		Head: "mulop",
		Body: "mod",
	},
	"66": {
		Head: "mulop",
		Body: "and",
	},
	"67": {
		Head: "factor",
		Body: "id",
	},
	"68": {
		Head: "factor",
		Body: "id ( expr_list )",
	},
	"69": {
		Head: "factor",
		Body: "int",
	},
	"70": {
		Head: "factor",
		Body: "float",
	},
	"71": {
		Head: "factor",
		Body: "string",
	},
	"72": {
		Head: "factor",
		Body: "boolean",
	},
	"73": {
		Head: "factor",
		Body: "( expr )",
	},
	"74": {
		Head: "factor",
		Body: "not factor",
	},
	"75": {
		Head: "sign",
		Body: "+",
	},
	"76": {
		Head: "sign",
		Body: "-",
	},
}
