package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Type string

const Symbol Type = "SYMBOL"
const Number Type = "NUMBER"

type token struct {
	Type   Type
	Number int
	Symbol rune
}

func tokenise(s string) []token {
	s = strings.ReplaceAll(s, "(", " ( ")
	s = strings.ReplaceAll(s, ")", " ) ")
	tokens := strings.Split(s, " ")
	var result []token
	for _, t := range tokens {
		t := strings.TrimSpace(t)
		if t != "" {
			switch t {
			case "(":
				result = append(result, token{Type: Symbol, Symbol: '('})
			case ")":
				result = append(result, token{Type: Symbol, Symbol: ')'})
			case "+":
				result = append(result, token{Type: Symbol, Symbol: '+'})
			case "*":
				result = append(result, token{Type: Symbol, Symbol: '*'})
			default:
				n, err := strconv.Atoi(t)
				if err != nil {
					panic(err)
				}
				result = append(result, token{Type: Number, Number: n})
			}
		}
	}
	return result
}

func containsBracket(tokens []token) *int {
	for i, v := range tokens {
		if v.Type == Symbol && v.Symbol == '(' {
			return &i
		}
	}
	return nil
}

func containsAddition(tokens []token) *int {
	for i, v := range tokens {
		if v.Type == Symbol && v.Symbol == '+' {
			return &i
		}
	}
	return nil
}

func closingBracket(tokens []token, i int) int {
	innerBrackets := 0
	for j, v := range tokens {
		if j > i {
			if v.Type == Symbol && v.Symbol == '(' {
				innerBrackets++
			}
			if v.Type == Symbol && v.Symbol == ')' {
				if innerBrackets == 0 {
					return j
				}
				innerBrackets--
			}
		}
	}
	panic("no closing bracket")
}

func parse(tokens []token) token {
	firstBracket := containsBracket(tokens)
	if firstBracket != nil {
		closeOfBracket := closingBracket(tokens, *firstBracket)
		next := tokens[:*firstBracket]
		next = append(next, parse(tokens[*firstBracket+1:closeOfBracket]))
		if closeOfBracket+2 < len(tokens) {
			next = append(next, tokens[closeOfBracket+1:]...)
		}
		return parse(next)
	}
	total := tokens[0].Number
	for i := 1; i < len(tokens); i += 2 {
		switch tokens[i].Symbol {
		case '+':
			total += tokens[i+1].Number
		case '*':
			total *= tokens[i+1].Number
		default:
			panic("unexpected")
		}
	}
	return token{Type: Number, Number: total}
}

func parse2(tokens []token) token {
	firstBracket := containsBracket(tokens)
	if firstBracket != nil {
		closeOfBracket := closingBracket(tokens, *firstBracket)
		next := tokens[:*firstBracket]
		next = append(next, parse2(tokens[*firstBracket+1:closeOfBracket]))
		if closeOfBracket+2 < len(tokens) {
			next = append(next, tokens[closeOfBracket+1:]...)
		}
		return parse2(next)
	}
	addition := containsAddition(tokens)
	if addition != nil {
		var next []token
		if *addition-1 >= 0 {
			next = append(next, tokens[:*addition-1]...)
		}
		next = append(next, token{Type: Number, Number: tokens[*addition-1].Number + tokens[*addition+1].Number})
		if *addition+2 < len(tokens) {
			next = append(next, tokens[*addition+2:]...)
		}
		return parse2(next)
	}
	total := tokens[0].Number
	for i := 1; i < len(tokens); i += 2 {
		switch tokens[i].Symbol {
		case '*':
			total *= tokens[i+1].Number
		default:
			panic("unexpected")
		}
	}
	return token{Type: Number, Number: total}
}

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	result1 := 0
	result2 := 0
	for _, line := range lines {
		tokens := tokenise(line)
		result1 += parse(tokens).Number
		tokens = tokenise(line)
		result2 += parse2(tokens).Number
	}
	fmt.Printf("result1 = %d\n", result1)
	fmt.Printf("result2 = %d\n", result2)

}
