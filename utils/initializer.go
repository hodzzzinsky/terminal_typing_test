package utils

import (
	"bufio"
	"log"
	"math/rand/v2"
	"os"
)

const TASK_SIZE = 250

func ReadFromFile() []string {

	array := []string{}

	file, err := os.Open("./utils/words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		word := scan.Text()
		array = append(array, word)
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
	// todo: alocate array
	return result
}
