package lexer

import (
	"testing"

	"github.com/kianooshaz/paca/tokens"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		source string
		tokens []tokens.Token
	}{
		// {
		// 	desc:   "1: hello world",
		// 	source: `"hello world"`,
		// 	tokens: []tokens.Token{{Type: tokens.STRING, Value: "hello world"}},
		// },
		// {
		// 	desc:   "2: string with numbers",
		// 	source: `"hello world" 10`,
		// 	tokens: []tokens.Token{
		// 		{Type: tokens.STRING, Value: "hello world"},
		// 		{Type: tokens.INTEGER, Value: "10"},
		// 	},
		// },
		// {
		// 	desc:   "3: string with backslash n",
		// 	source: `"hello \n world"`,
		// 	tokens: []tokens.Token{
		// 		{Type: tokens.STRING, Value: `hello \n world`},
		// 	},
		// },
		// {
		// 	desc:   "4: string with comment",
		// 	source: `"hello world" // 303`,
		// 	tokens: []tokens.Token{
		// 		{Type: tokens.STRING, Value: "hello world"},
		// 	},
		// },
		{
			desc: "5: string with comment and new line",
			source: `"hello world" // 303
			"hello future"`,
			tokens: []tokens.Token{
				{Type: tokens.STRING, Value: "hello world"},
				{Type: tokens.STRING, Value: "hello future"},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := New(tC.source)
			for i, r := range tC.tokens {
				if r != l.tokens[i] {
					t.Fail()
					t.Logf("want: %v ,\t but get : %v", r, tC.tokens[i])
				}
			}
		})
	}
}
