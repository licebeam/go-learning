package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const str = "สวัสดี" // had to copy this one lol
	fmt.Println("len", len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}

	fmt.Println()
	fmt.Println("rune count:", utf8.RuneCountInString(str))

	for idx, runeValue := range str {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found rune t")
	} else if r == 'ส' {
		fmt.Println("Found so sua")
	}
}
