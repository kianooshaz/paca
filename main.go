package main

import (
	"log"
	"os"
	"strings"

	"github.com/kianooshaz/paca/lexer"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	str := string(input)
	str = strings.NewReplacer(`\n`, "\n", `\t`, "\t", `\r`, "\r").Replace(str)
	l := lexer.New(str)
	l.PrintTokens()
}
