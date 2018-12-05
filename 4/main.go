package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

const example = "pqrstuv"
const input = "iwrupvqb"

func main() {
	fmt.Println(puzzle(input, "00000"))
	fmt.Println(puzzle(input, "000000"))
}

func puzzle(input, prefix string) int {
	for i := 0; ; i++ {
		secret := fmt.Sprintf("%s%d", input, i)

		result := fmt.Sprintf("%x", md5.Sum([]byte(secret)))

		if strings.HasPrefix(result, prefix) {
			return i
		}
	}
}
