package gui

import (
	"bowling-game/bowling"
	"fmt"
)

func PrintFrames(frames []*bowling.Frame) {
	fmt.Println("----  Bowling-Results  ----")
	for _, currFrame := range frames {
		printLine1(currFrame)
	}
	fmt.Println()

	for _, currFrame := range frames {
		printLine3(currFrame)
	}
	fmt.Println()
}

func printLine1(frame *bowling.Frame) {
	var firstScore string
	var secondScore string
	if frame.KnockDowns[0] == bowling.NumberOfPins {
		firstScore = " X"
		secondScore = "  "
	} else {
		firstScore = fmt.Sprintf(" %d", frame.KnockDowns[0])
		if frame.KnockDowns[1] == bowling.NumberOfPins {
			secondScore = " X"
		} else if frame.HadSpare() {
			secondScore = " /"
		} else {
			secondScore = fmt.Sprintf(" %d", frame.KnockDowns[1])
		}
	}

	fmt.Printf("%s %s | ", firstScore, secondScore)
}

func printLine3(frame *bowling.Frame) {
	if frame.Points < 10 {
		fmt.Print(" ")
	}
	fmt.Printf("  %d    ", frame.Points)
}
