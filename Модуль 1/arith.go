package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Lexem struct {
	Tag
	Image string
}

type Tag int

const (
	ERROR  Tag = 1 << iota // Неправильная лексема
	NUMBER                 // Целое число
	VAR                    // Имя переменной
	PLUS                   // Знак +
	MINUS                  // Знак -
	MUL                    // Знак *
	DIV                    // Знак /
	LPAREN                 // Левая круглая скобка
	RPAREN                 // Правая круглая скобка
)

func writeIn(word, digit *[]int32, lexems chan Lexem) {
	if len(*word) > 0 {
		lexems <- Lexem{VAR, string(*word)}
		*word = (*word)[:0]
	} else if len(*digit) > 0 {
		lexems <- Lexem{NUMBER, string(*digit)}
		*digit = (*digit)[:0]
	}
}

func lexer(expr string, lexems chan Lexem) {
	word := make([]int32, 0, len(expr))
	digit := make([]int32, 0, len(expr))
	for _, x := range expr {
		if x >= '0' && x <= '9' {
			if len(word) > 0 {
				word = append(word, x)
			} else {
				digit = append(digit, x)
			}
		} else if x >= 'a' && x <= 'z' || x >= 'A' && x <= 'Z' {
			if len(digit) == 0 {
				word = append(word, x)
			}
		} else {
			writeIn(&word, &digit, lexems)
			switch x {
			case '+':
				lexems <- Lexem{PLUS, "+"}
			case '-':
				lexems <- Lexem{MINUS, "-"}
			case '*':
				lexems <- Lexem{MUL, "*"}
			case '/':
				lexems <- Lexem{DIV, "/"}
			case '(':
				lexems <- Lexem{LPAREN, "("}
			case ')':
				lexems <- Lexem{RPAREN, ")"}
			case ' ':
			default:
				lexems <- Lexem{ERROR, string(x)}
			}
		}
	}
	writeIn(&word, &digit, lexems)
	close(lexems)
}

func eval_(lexems []Lexem) (bool, Lexem) {
	if lexems[0].Tag == MINUS {
		if len(lexems) == 2 {
			num, _ := strconv.Atoi(lexems[1].Image)
			return true, Lexem{NUMBER, strconv.Itoa(-num)}
		}
	} else if len(lexems)%2 == 1 {
		num, _ := strconv.Atoi(lexems[0].Image)
		var num2 int
		for i := 1; i < len(lexems); i += 2 {
			num2, _ = strconv.Atoi(lexems[i+1].Image)
			switch lexems[i].Tag {
			case PLUS:
				num += num2
			case MINUS:
				num -= num2
			case MUL:
				num *= num2
			case DIV:
				if num2 != 0 {
					num /= num2
				}
			default:
				return false, Lexem{}
			}
		}
		return true, Lexem{NUMBER, strconv.Itoa(num)}
	}
	return false, Lexem{}
}

func E(lexems []Lexem) (bool, Lexem) {
	var (
		a, b bool
		x    Lexem
		y    []Lexem
	)
	for i := 1; i <= len(lexems); i++ {
		a, x = T(lexems[:i])
		b, y = E_(lexems[i:])
		if a && b {
			return eval_(append([]Lexem{x}, y...))
		}
	}
	return false, Lexem{}
}

func E_(lexems []Lexem) (bool, []Lexem) {
	if len(lexems) == 0 {
		return true, []Lexem{}
	}
	if lexems[0].Tag&(PLUS|MINUS) != 0 {
		var (
			a, b bool
			x    Lexem
			y    []Lexem
		)
		for i := 2; i <= len(lexems); i++ {
			a, x = T(lexems[1:i])
			b, y = E_(lexems[i:])
			if a && b {
				return true, append([]Lexem{lexems[0], x}, y...)
			}
		}
	}
	return false, []Lexem{}
}

func T(lexems []Lexem) (bool, Lexem) {
	var (
		a, b bool
		x    Lexem
		y    []Lexem
	)
	for i := 1; i <= len(lexems); i++ {
		a, x = F(lexems[:i])
		b, y = T_(lexems[i:])
		if a && b {
			return eval_(append([]Lexem{x}, y...))
		}
	}
	return false, Lexem{}
}

func T_(lexems []Lexem) (bool, []Lexem) {
	if len(lexems) == 0 {
		return true, []Lexem{}
	}
	if lexems[0].Tag&(MUL|DIV) != 0 {
		var (
			a, b bool
			x    Lexem
			y    []Lexem
		)
		for i := 2; i <= len(lexems); i++ {
			a, x = F(lexems[1:i])
			b, y = T_(lexems[i:])
			if a && b {
				return true, append([]Lexem{lexems[0], x}, y...)
			}
		}
	}
	return false, []Lexem{}
}

func F(lexems []Lexem) (bool, Lexem) {
	if len(lexems) == 1 {
		if lexems[0].Tag&(VAR|NUMBER) != 0 {
			return true, lexems[0]
		}
	} else if lexems[0].Tag == LPAREN {
		if lexems[len(lexems)-1].Tag == RPAREN {
			return E(lexems[1 : len(lexems)-1])
		}
	} else if lexems[0].Tag == MINUS {
		if a, x := F(lexems[1:]); a {
			return eval_([]Lexem{lexems[0], x})
		}
	}
	return false, Lexem{}
}

var (
	varIn  = make(map[string]bool)
	varVal = make(map[string]string)
)

func main() {
	sentence, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	lexems := make(chan Lexem)
	go lexer(sentence[:len(sentence)-2], lexems)

	lexemsSyntax := make([]Lexem, 0, 20)
	lexemsEval := lexemsSyntax
	for x := range lexems {
		if x.Tag == ERROR {
			fmt.Print("error")
			return
		} else if x.Tag == VAR {
			lexemsSyntax = append(lexemsSyntax, Lexem{NUMBER, "0"})
		} else {
			lexemsSyntax = append(lexemsSyntax, x)
		}
		lexemsEval = append(lexemsEval, x)
	}

	noErr, n := E(lexemsSyntax)
	if noErr {
		for i, x := range lexemsEval {
			if x.Tag == VAR {
				if !varIn[x.Image] {
					var val int
					fmt.Printf("Enter the value of var %s: ", x.Image)
					fmt.Scan(&val)
					varIn[x.Image] = true
					varVal[x.Image] = strconv.Itoa(val)
				}
				lexemsEval[i] = Lexem{NUMBER, varVal[x.Image]}
			}
		}
		_, n = E(lexemsEval)
		fmt.Print(n.Image)
	} else {
		fmt.Print("error")
	}
}
