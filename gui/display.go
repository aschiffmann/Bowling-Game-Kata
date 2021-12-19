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
	fmt.Printf("|%d|%d| ", frame.KnockDowns[0], frame.KnockDowns[1])
}

func printLine3(frame *bowling.Frame) {
	fmt.Printf(" %d   ", frame.Points)
}
