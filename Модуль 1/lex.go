package main

import "fmt"

func lex(sentence string, array AssocArray) []int {
	lenSentence := len(sentence)
	var (
		word   string
		n      int
		exists bool
	)
	a := 0
	res := make([]int, 0, lenSentence>>2+1)
	ind := make([]uint8, 0, lenSentence)
	for i := 0; i <= lenSentence; i++ {
		if i == lenSentence || sentence[i] == ' ' {
			if len(ind) > 0 {
				word = string(ind)
				n, exists = array.Lookup(word)
				if exists {
					res = append(res, n)
				} else {
					a++
					array.Assign(word, a)
					res = append(res, a)
				}
				ind = ind[:0]
			}
		} else {
			ind = append(ind, sentence[i])
		}
	}
	return res
}

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

type Array map[string]int

func (a Array) Assign(s string, x int) {
	a[s] = x
}

func (a Array) Lookup(s string) (x int, exists bool) {
	x = a[s]
	exists = x != 0
	return
}

func main() {
	sentence := "alpha x1 beta alpha x1 y"
	for _, x := range lex(sentence, Array(make(map[string]int))) {
		fmt.Printf("%d ", x)
	}
}
