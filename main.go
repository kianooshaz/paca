package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kianooshaz/paca/lexer"
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
	str := string(input)
	str = strings.NewReplacer(`\n`, "\n", `\t`, "\t", `\r`, "\r").Replace(str)
	it := tokens.NewIdentTable()
	l := lexer.New(str, it)

	fmt.Println("########## idents table ##########")
	l.PrintIdentTable()
	fmt.Println("########## tokens ##########")
	l.PrintTokens()
}
