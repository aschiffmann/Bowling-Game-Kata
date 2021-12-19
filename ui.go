package main

import (
	"fmt"
)

const strikeDisplay = "X"
const spareDisplay = "/"
const noScoreDisplay = "-"

func printResults(frames []*frame) {
	fmt.Println("----  Bowling-Results  ----")
	for _, f := range frames {
		printScoresLine(f)
	}
	fmt.Println()

	for _, f := range frames {
		printTotalPointsLine(f)
	}
}

func printScoresLine(f *frame) {
	firstScore := getScoreDisplay(f.knockDowns[0])

	var secondScore string
	if f.hadSpare() {
		secondScore = spareDisplay
	} else {
		secondScore = getScoreDisplay(f.knockDowns[1])
	}

	var thirdScore string
	if len(f.knockDowns) == 3 {
		if !f.hadSpare() && f.points == totalPins {
			thirdScore = fmt.Sprintf(" %s", spareDisplay)
		} else {
			thirdScore = fmt.Sprintf(" %s", getScoreDisplay(f.knockDowns[2]))
		}
	}

	fmt.Printf("%s  %s%s | ", firstScore, secondScore, thirdScore)
}

func printTotalPointsLine(frame *frame) {
	totalPoints := frame.getTotalPoints()
	if totalPoints < 10 {
		fmt.Print(" ")
	}
	fmt.Printf(" %d   ", totalPoints)

	if totalPoints < 100 {
		fmt.Printf(" ")
	}
}

func getScoreDisplay(score int) string {
	switch score {
	case -1:
		return noScoreDisplay
	case totalPins:
		return strikeDisplay
	default:
		return fmt.Sprintf("%d", score)
	}
}
