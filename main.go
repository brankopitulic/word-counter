package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	str := string(data)

	dataScanner := bufio.NewScanner(strings.NewReader(str))
	dataScanner.Split(bufio.ScanWords)
	count := 0
	for dataScanner.Scan() {
		count++
	}

	fmt.Println("Total characters count =", utf8.RuneCountInString(str))
	fmt.Println("Total words count =", count)
	fmt.Println("Total lines count =", countRune(str, '\n'))
}

func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count + 1 // +1 count last line
}
