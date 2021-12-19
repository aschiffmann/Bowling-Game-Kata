package main

import (
	"bowling-game/bowling"
	"bowling-game/gui"
	"fmt"
)

const NumberOfFrames = 10
const NumberOfPins = 10

func main() {
	fmt.Println("hello to bowling!")

	var playedFrames []*bowling.Frame

	var previousFrame *bowling.Frame
	for currFrameNr := 1; currFrameNr <= NumberOfFrames; currFrameNr++ {
		currFrame := bowling.Frame{PreviousFrame: previousFrame, FrameNr: currFrameNr}
		currFrame.ExecuteRolls()
		playedFrames = append(playedFrames, &currFrame)
		previousFrame = &currFrame
	}

	gui.PrintFrames(playedFrames)
}
