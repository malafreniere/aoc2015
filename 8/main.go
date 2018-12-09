package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	// b, err := ioutil.ReadFile("example.txt")

	if err != nil {
		log.Fatalf("%v", err)
	}

	content := string(b)
	words := strings.Split(content, "\r\n")

	c, m := puzzle1(words)

	fmt.Printf("%d\n", c-m)

	e := puzzle2(words)

	fmt.Printf("%d\n", e-c)
}

func puzzle1(words []string) (codeChars, memoryChars int) {
	codeChars = 0
	memoryChars = 0

	for _, w := range words {
		codeChars += len(w)

		for i := 1; i < len(w)-1; i++ {
			memoryChars++

			if w[i] == '\\' {
				i++

				if w[i] == 'x' {
					i += 2
				}
			}
		}
	}

	return codeChars, memoryChars
}

func puzzle2(words []string) int {
	count := 0

	for _, w := range words {
		encoded := strings.Builder{}
		encoded.WriteRune('"')
		for i := 0; i < len(w); i++ {
			if w[i] == '"' {
				encoded.WriteString("\\\"")
			} else if w[i] == '\\' {
				encoded.WriteString("\\\\")
			} else {
				encoded.WriteByte(w[i])
			}
		}
		encoded.WriteRune('"')
		count += encoded.Len()
	}

	return count
}
