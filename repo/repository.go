package repo

import (
	"math"
	"strings"
	"ttt/utils"
)

var Task []string = []string{}
var result []string = []string{}
var Index = 0

var Colored []string = []string{}

const PORTION_SIZE = 5

const (
	CYAN      = "\033[0;36m"
	GREEN     = "\033[0;32m"
	RED       = "\033[0;31m"
	CYAN_BOLD = "\033[1;36m"
	RED_BOLD  = "\033[1;31m"
)

type Statistics struct {
	WPM             float64
	Accuracy        float64
	All_letters     int
	Wrong_letters   int
	Correct_letters int
	Correct_words   int
}

func Append(char rune) {
	if char == '\x00' {
		Index++
	} else if len(result) == 0 {
		result = append(result, string(char))
	} else {
		if Index >= len(result) {
			result = append(result, string(char))
		} else {
			result[Index] += string(char)
		}
	}
}

func DeleteLast() {
	if len(result) > 0 && len(result) > Index {
		if len(result[Index]) > 0 {
			var word = result[Index]
			var without = word[:len(word)-1]
			result[Index] = without
		}
	}
}

func InitTask() {
	Task = utils.ReadFromFile()
	for _, w := range Task {
		Colored = append(Colored, CYAN+w)
	}
}

func GetCurrentWord() string {
	if Index < len(result) {
		return result[Index]
	}
	return ""
}

func isFinished(indx int) bool {
	if len(result) <= indx {
		return false
	}
	rw := strings.TrimSpace(result[indx])
	tw := strings.TrimSpace(Task[indx])

	return len(rw) == len(tw)
}

func isCharCorrect(indx int) bool {
	if len(result) <= indx {
		return true
	}
	rw := strings.TrimSpace(result[indx])
	tw := strings.TrimSpace(Task[indx])
	if len(rw) > len(tw) {
		return false
	}

	correct := true

	for i := 0; i < len(rw); i++ {
		if rw[i] != tw[i] {
			correct = false
		}
	}
	return correct
}

func ColorArray() {
	if isCharCorrect(Index) && !isFinished(Index) {
		colorCorrect(Index)
	} else if !isCharCorrect(Index) && !isFinished(Index) {
		colorIncorrect(Index)
	} else if isCharCorrect(Index) && isFinished(Index) {
		colorCorrectFin(Index)
	} else {
		colorIncorrectFin(Index)
	}
	if Index > 0 {
		prev := Index - 1
		if isCharCorrect(prev) && isFinished(prev) {
			colorCorrectFin(prev)
		} else {
			colorIncorrectFin(prev)
		}
	}
}

func colorCorrect(indx int) {
	Colored[indx] = CYAN_BOLD + Task[indx]
}

func colorIncorrect(indx int) {
	Colored[indx] = RED_BOLD + Task[indx]
}

func colorCorrectFin(indx int) {
	Colored[indx] = GREEN + Task[indx]
}

func colorIncorrectFin(indx int) {
	Colored[indx] = RED + Task[indx]
}

func CalculateStatistics() Statistics {

	wrong_letters := 0.0
	all_letters := 0.0

	for i := 0; i < len(result); i++ {
		tw := Task[i]
		rw := result[i]

		if len(rw) != len(tw) {
			if len(rw) > len(tw) {
				wrong_letters += float64(len(tw))
			} else {
				for j := 0; j < len(rw); j++ {
					if rw[j] != tw[j] {
						wrong_letters += 1
					}
				}
				wrong_letters += float64(len(tw) - len(rw))
			}
		} else {
			for j := 0; j < len(tw); j++ {
				if rw[j] != tw[j] {
					wrong_letters += 1
				}
			}
		}
		all_letters += float64(len(tw))
	}

	avg := calcAvgWordSize(all_letters)
	netWPM := calcNetWPM(all_letters, wrong_letters, avg)
	wpm := netWPM * 100
	acc := acc2((all_letters - wrong_letters), all_letters)
	corr := countCorrectWords()

	stats := Statistics{
		wpm,
		acc,
		int(all_letters),
		int(wrong_letters),
		int(all_letters - wrong_letters),
		corr,
	}
	return stats
}

func calcNetWPM(all float64, wrong float64, avg_word_len int) float64 {
	avg := all / float64(avg_word_len)
	cor := math.Abs(avg - wrong)
	return (cor / 60)
}

func acc2(cor float64, all float64) float64 {
	return (cor / all) * 100
}

func countCorrectWords() int {
	corr := 0
	for i := 0; i < len(result); i++ {
		rw := result[i]
		tw := Task[i]

		if tw == rw {
			corr += 1
		}
	}
	return corr
}

func calcAvgWordSize(all_letters float64) int {
	avg := all_letters / float64(len(result))
	return int(math.Round(avg))
}
