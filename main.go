package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kianooshaz/paca/lexer"
	"github.com/kianooshaz/paca/parser"
	"github.com/kianooshaz/paca/tokens"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	sourceCode := string(input)
	sourceCode = strings.NewReplacer(`\n`, "\n", `\t`, "\t", `\r`, "\r").Replace(sourceCode)
	it := tokens.NewIdentTable()
	l := lexer.New(sourceCode, it)

	actionTableFile, _ := os.ReadFile("./grammar/data/actionTable.json")
	var actionTable parser.Table
	json.Unmarshal(actionTableFile, &actionTable)

	gotoTableFile, _ := os.ReadFile("./grammar/data/gotoTable.json")
	var gotoTable parser.Table
	json.Unmarshal(gotoTableFile, &gotoTable)

	fmt.Println("########## idents table ##########")
	l.PrintIdentTable()
	fmt.Println("########## tokens ##########")
	l.PrintTokens()

	p := parser.New(actionTable, gotoTable)
	p.Parse(l)
	fmt.Println("########## right-most derivation ##########")
	p.PrintDerivation()
}
