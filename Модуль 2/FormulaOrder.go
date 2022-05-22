package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func uploadVar(varsCount *int, varsId *map[string]int, partGraph *[]map[int]bool, variable *string, equal bool, waitComma, waitOperator, value, flag *bool) bool {
	if *value {
		if *waitOperator {
			return true
		}
		*waitOperator = true
		*value = false
		*flag = true
	} else if len(*variable) != 0 {
		if !equal && *waitComma || equal && *waitOperator {
			return true
		}
		if (*varsId)[*variable] == 0 {
			(*varsId)[*variable] = *varsCount
			*varsCount++
		}
		if equal {
			(*partGraph)[1][(*varsId)[*variable]] = true
		} else {
			(*partGraph)[0][(*varsId)[*variable]] = true
		}
		*variable = ""
		if equal {
			*waitOperator = true
		} else {
			*waitComma = true
		}
		*flag = true
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	newLine, _ := in.ReadString('\n')

	sentencesCount := 0
	varsCount := 0
	varsId := map[string]int{}
	var (
		graph                                       [][]map[int]bool
		sentences                                   []string
		equal, value, waitOperator, waitComma, flag bool
		variable                                    string
		i, countArgs, brackets                      int
		char                                        rune
	)

	for len(newLine) != 2 {
		graph = append(graph, []map[int]bool{{}, {}})
		sentences = append(sentences, newLine)
		equal, value, waitComma, waitOperator, flag = false, false, false, false, false
		variable = ""
		brackets = 0
		for i = range newLine {
			char = rune(newLine[i])
			if unicode.IsDigit(char) {
				if len(variable) > 0 {
					variable += string(char)
				} else if equal && !waitOperator {
					value = true
				} else {
					fmt.Println("syntax error")
					return
				}
			} else if unicode.IsLetter(char) {
				if value || !equal && waitComma || equal && waitOperator {
					fmt.Println("syntax error")
					return
				} else {
					variable += string(char)
				}
			} else {
				if uploadVar(&varsCount, &varsId, &graph[sentencesCount], &variable, equal, &waitComma, &waitOperator, &value, &flag) {
					fmt.Println("syntax error")
					return
				}
				if char == ',' {
					if equal {
						if waitOperator && countArgs != 0 && flag {
							countArgs--
							waitOperator = false
						} else {
							fmt.Println("syntax error")
							return
						}
						flag = false
					} else if waitComma {
						countArgs++
						waitComma = false
					} else {
						fmt.Println("syntax error")
						return
					}
				} else if char == '=' {
					if equal || !waitComma {
						fmt.Println("syntax error")
						return
					}
					equal = true
					waitOperator = false
					flag = false
				} else if char == '(' {
					if equal && !waitOperator {
						brackets++
					} else {
						fmt.Println("syntax error")
						return
					}
				} else if char == ')' {
					if equal && waitOperator && brackets != 0 {
						brackets--
					} else {
						fmt.Println("syntax error")
						return
					}
				} else if char == '+' || char == '-' || char == '*' || char == '/' {
					if equal && waitOperator {
						waitOperator = false
					} else {
						fmt.Println("syntax error")
						return
					}
				}
			}
		}
		if uploadVar(&varsCount, &varsId, &graph[sentencesCount], &variable, equal, &waitComma, &waitOperator, &value, &flag) ||
			!equal || !waitOperator || !flag || countArgs != 0 || brackets != 0 {
			fmt.Println("syntax error")
			return
		}
		sentencesCount++
		newLine, _ = in.ReadString('\n')
	}

	result := make([]int, sentencesCount)
	queue := map[int]bool{}
	for i = 0; i < sentencesCount; i++ {
		queue[i] = true
	}
	var a, b int
	for j := 0; j < sentencesCount; j++ {
		flag = true
		for i = range queue {
			if len(graph[i][1]) == 0 {
				delete(queue, i)
				flag = false
				result[j] = i
				for a = range queue {
					for b = range graph[i][0] {
						if graph[a][1][b] {
							delete(graph[a][1], b)
						}
					}
				}
				break
			}
		}
		if flag {
			fmt.Println("«cycle»")
			return
		}
	}
	for _, i = range result {
		fmt.Print(sentences[i])
	}
}
