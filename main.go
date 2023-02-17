package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordList() []string {

	f, err := os.Open("word-list")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	var r []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r = append(r, scanner.Text())
	}

	return r
}

func input() string {
	// get input from terminal
	r := os.Args[1]
	// remove redundant spaces
	r = strings.Join(strings.Fields(r), " ")
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

func lower(word string) string {
	var arr []string
	for _, w := range word {
		arr = append(arr, string(w))
	}
	arr[0] = strings.ToLower(arr[0])
	r := strings.Join(arr, "")
	return r
}

func ignore(s string) bool {
	// check if word is within ignore list for title case
	list := wordList()
	for _, wl := range list {
		if s == wl {
			return true
		}
	}
	return false
}

func titleCase(word []string) string {
	for idx, w := range word {
		if idx == 0 {
			word[idx] = capital(w)
		}
		if ignore(w) != true && idx != 0 {
			word[idx] = capital(w)
		}
	}
	r := strings.Join(word, " ")
	return r
}

func main() {
	input := input()
	words := words(input)
	fmt.Println(titleCase(words))
}
