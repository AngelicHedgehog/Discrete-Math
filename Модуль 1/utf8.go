package main

import "fmt"

func encode(utf32 []rune) []byte {
	utf8 := make([]byte, 0, len(utf32)<<2)
	for _, x := range utf32 {
		if x>>7 == 0 {
			utf8 = append(utf8, byte(x))
		} else if x>>11 == 0 {
			utf8 = append(utf8, byte(3<<6+x>>6), byte(1<<7+x%(1<<6)))
		} else if x>>16 == 0 {
			utf8 = append(utf8,
				byte(7<<5+x>>12),
				byte(1<<7+x>>6%(1<<6)),
				byte(1<<7+x%(1<<6)))
		} else {
			utf8 = append(utf8,
				byte(15<<4+x>>18),
				byte(1<<7+x>>12%(1<<6)),
				byte(1<<7+x>>6%(1<<6)),
				byte(1<<7+x%(1<<6)))
		}
	}
	return utf8
}

func decode(utf8 []byte) []rune {
	utf32 := make([]rune, 0, len(utf8))
	for i := 0; i < len(utf8); i++ {
		if utf8[i]>>7 == 0 {
			utf32 = append(utf32, rune(utf8[i]))
		} else if len(utf8) >= 2 && utf8[i]>>5 == 6 && utf8[i+1]>>6 == 2 {
			utf32 = append(utf32, rune(utf8[i])%(1<<5)<<6+rune(utf8[i+1])%(1<<6))
			i++
		} else if len(utf8) >= 3 && utf8[i]>>4 == 14 &&
			utf8[i+1]>>6 == 2 &&
			utf8[i+2]>>6 == 2 {
			utf32 = append(utf32, rune(utf8[i])%(1<<4)<<12+
				rune(utf8[i+1])%(1<<6)<<6+
				rune(utf8[i+2])%(1<<6))
			i += 2
		} else if len(utf8) >= 4 && utf8[i]>>3 == 30 &&
			utf8[i+1]>>6 == 2 &&
			utf8[i+2]>>6 == 2 &&
			utf8[i+3]>>6 == 2 {
			utf32 = append(utf32, rune(utf8[i])%(1<<3)<<18+
				rune(utf8[i+1])%(1<<6)<<12+
				rune(utf8[i+2])%(1<<6)<<6+
				rune(utf8[i+3])%(1<<6))
			i += 3
		}
	}
	return utf32
}

func main() {
	fmt.Printf("%s", string(decode(encode([]rune("When bot will done? Или, если по-русски, КОГДА УЖЕ БОТ БУДЕТ?")))))
}
