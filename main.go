package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
	"time"
	"ttt/repo"
)

var general []string = []string{}
var preview []string = []string{}
var startTime time.Time
var started = false

func main() {
	repo.InitTask()
	go drawUI()
	processInput()
}

func processInput() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		//replace with switch case
		if key == keyboard.KeyEsc {
			drawStats()
			os.Exit(0)
		} else if key == keyboard.KeyBackspace ||
			key == keyboard.KeyDelete ||
			key == keyboard.KeyBackspace2 {
			repo.DeleteLast()
		} else {
			repo.Append(char)
			//fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		}
		if !started {
			startTime = time.Now()
			started = true
		}
		repo.ColorArray()
	}
}

func showTimer() float64 {
	diff := time.Now().Sub(startTime)
	left := 60 - diff.Seconds()
	if left < -1 {
		left = 60
	}
	return left
}

func sliceArray() {
	if repo.Index%repo.PORTION_SIZE == 0 {
		general = repo.Colored[repo.Index:(repo.Index + repo.PORTION_SIZE)]
		preview = repo.Colored[(repo.Index + repo.PORTION_SIZE):(repo.Index + (repo.PORTION_SIZE * 2))]
	}
}

func drawArray() {
	sliceArray()
	for _, w := range general {
		fmt.Print(w + " ")
	}
	fmt.Println()

	for _, w := range preview {
		fmt.Print(w + " ")
	}
	fmt.Println()
}

func drawStats() {
	stats := repo.CalculateStatistics()
	fmt.Println("\033[H\033[2J")
	fmt.Printf("wpm: %.0f \n", stats.WPM)
	fmt.Printf("accuracy: %.0f", stats.Accuracy)
	fmt.Print("% \n")
	fmt.Printf("all symbols: %d\n", stats.All_letters)
	fmt.Printf("corret symols: %d\n", stats.Correct_letters)
	fmt.Printf("wrong symbols: %d\n", stats.Wrong_letters)
	fmt.Printf("correct words: %d\n", stats.Correct_words)
}

func drawUI() {
	for {
		fmt.Println("\033[H\033[2J")
		fmt.Printf("press ESC to quit \n")
		fmt.Println()
		fmt.Printf("⏱️ %.0f", showTimer())
		fmt.Println()
		drawArray()
		fmt.Printf(">%s", repo.GetCurrentWord())
		time.Sleep(50 * time.Millisecond)
		if showTimer() <= 0 {
			drawStats()
			os.Exit(0)
		}
	}
}
