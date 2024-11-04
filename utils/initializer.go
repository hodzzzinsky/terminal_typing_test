package utils

import (
	"bufio"
	_ "embed"
	"log"
	"math/rand/v2"
	"strings"
)

const TASK_SIZE = 250

//go:embed words.txt
var words_file string

func ReadFromFile() []string {

	array := []string{}
	scan := bufio.NewScanner(strings.NewReader(words_file))

	for scan.Scan() {
		word := scan.Text()
		array = append(array, word)
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}
	return ReorderArray(array)
}

func ReorderArray(array []string) []string {
	length := len(array)
	result := []string{}

	for i := 0; i < TASK_SIZE; i++ {
		word := array[rand.IntN(length)]
		result = append(result, word)
	}
	return result
}
