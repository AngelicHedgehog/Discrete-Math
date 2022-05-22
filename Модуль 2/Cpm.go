package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"unicode"
)

func uploadWord(vals *map[int]int, wordToId *map[string]int, idToWord *map[int]string, graph *map[int][]map[int]bool, lastWord int, num int, word string) {
	var id int
	if (*wordToId)[word] == 0 {
		id = len(*wordToId) + 1
		(*wordToId)[word] = id
		(*idToWord)[id] = word
		(*vals)[id] = num
		(*graph)[id] = []map[int]bool{{}, {}}
	} else {
		id = (*wordToId)[word]
	}
	if lastWord != 0 {
		(*graph)[lastWord][0][id] = true
		(*graph)[id][1][lastWord] = true
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	newLine, _ := in.ReadString('\n')

	vals := make(map[int]int)
	wordToId := map[string]int{}
	idToWord := map[int]string{}
	graph := map[int][]map[int]bool{}
	lastWord := 0
	word := ""
	isWord := false
	num := 0
	var (
		lastChar rune
		char     rune
	)
	for true {
		for i := range newLine {
			char = rune(newLine[i])
			if unicode.IsSpace(char) {
				continue
			} else {
				lastChar = char
				if unicode.IsLetter(char) {
					isWord = true
					word += string(char)
				} else if char == '(' {
					isWord = false
				} else if unicode.IsDigit(char) {
					if isWord {
						word += string(char)
					} else {
						num = num*10 + int(char-'0')
					}
				} else if newLine[i] == '<' || newLine[i] == ';' {
					isWord = false
					uploadWord(&vals, &wordToId, &idToWord, &graph, lastWord, num, word)
					if newLine[i] == '<' {
						lastWord = wordToId[word]
					} else {
						lastWord = 0
					}
					word = ""
					num = 0
				}
			}
		}
		if lastChar != ';' && lastChar != '<' {
			break
		}
		newLine, _ = in.ReadString('\n')
	}
	uploadWord(&vals, &wordToId, &idToWord, &graph, lastWord, num, word)

	var (
		a, b  int
		A, B  map[int]bool
		queue []int
	)
	blues := map[int]bool{}
	for i := range vals {
		if !blues[i] {
			queue = []int{i}
			A = map[int]bool{}
			for len(queue) != 0 {
				a = queue[0]
				queue = queue[1:]
				A[a] = true
				for b = range graph[a][0] {
					if !A[b] {
						queue = append(queue, b)
					}
				}
			}
			queue = []int{i}
			B = map[int]bool{}
			for len(queue) != 0 {
				a = queue[0]
				queue = queue[1:]
				B[a] = true
				for b = range graph[a][1] {
					if !B[b] {
						queue = append(queue, b)
					}
				}
			}
			for a = range A {
				if a != i && B[a] {
					for b = range A {
						blues[b] = true
					}
				}
			}
		}
	}

	reds := map[int]bool{}
	max := 0
	var (
		queueWay       [][]int
		queueN, queueI []int
		flag           bool
	)
	for i := range vals {
		if len(graph[i][1]) == 0 && !blues[i] {
			queueWay = [][]int{{i}}
			queueN = []int{vals[i]}
			queueI = []int{i}
			for len(queueI) != 0 {
				flag = true
				for b = range graph[queueI[0]][0] {
					if !blues[b] {
						flag = false
						queueWay = append(queueWay, append(queueWay[0], b))
						queueN = append(queueN, queueN[0]+vals[b])
						queueI = append(queueI, b)
					}
				}
				if flag && queueN[0] >= max {
					if queueN[0] > max {
						reds = map[int]bool{}
						max = queueN[0]
					}
					for _, a = range queueWay[0] {
						reds[a] = true
					}
				}
				queueWay = queueWay[1:]
				queueN = queueN[1:]
				queueI = queueI[1:]
			}
		}
	}

	var sortedNames []string
	for el := range wordToId {
		sortedNames = append(sortedNames, el)
	}
	sort.Strings(sortedNames)
	var sortedNums []int
	for _, el := range sortedNames {
		sortedNums = append(sortedNums, wordToId[el])
	}

	fmt.Println("digraph {")
	for _, a = range sortedNums {
		fmt.Printf("\t%s [label = \"%s(%d)\"", idToWord[a], idToWord[a], vals[a])
		if reds[a] {
			fmt.Printf(", color = red]\n")
			//		} else if blues[a] {
			//			fmt.Printf(", color = blue]\n")
		} else {
			fmt.Printf("]\n")
		}
	}
	for _, a = range sortedNums {
		for _, b = range sortedNums {
			if graph[a][0][b] {
				fmt.Printf("\t%s -> %s", idToWord[a], idToWord[b])
				if reds[a] && reds[b] {
					fmt.Println(" [color = red]")
					//				} else if blues[a] && blues[b] {
					//					fmt.Println(" [color = blue]")
				} else {
					fmt.Println()
				}
			}
		}
	}
	fmt.Println("}")
}
