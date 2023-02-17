package main

import (
	"fmt"
	"os"
	"strings"
)

func input() string {
	r := os.Args[1]
	return r
}

func words(s string) []string {
	r := strings.Split(s, " ")
	return r
}

func capital(word string) string {
	var arr []string
	for _, w := range word {
		arr = append(arr, string(w))
	}
	arr[0] = strings.ToUpper(arr[0])
	r := strings.Join(arr, "")
	return r
}

func ignore(s string) bool {
	// check if word is within ignore list for title case
	return false
}

func titleCase(word []string) string {
	for idx, w := range word {
		word[idx] = capital(w)
	}
	r := strings.Join(word, " ")
	return r
}

func main() {
	input := input()
	words := words(input)

	fmt.Println(titleCase(words))
}
