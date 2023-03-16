package examples

import (
	"fmt"
	"unicode/utf8"
)

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'n' {
		fmt.Println("found sua")
	}
}

func StringsNRunes() {
	const s = "สวัสดี"

	fmt.Printf("Length: %v", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])

		if i == len(s) {
			fmt.Println()
		}
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeVal := range s {
		fmt.Printf("%#U starts at %d\n", runeVal, idx)
	}

	fmt.Println("\n Using DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}
