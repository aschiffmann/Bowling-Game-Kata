package main

import (
	"bowling-game/bowling"
	"bowling-game/gui"
	"math/rand"
	"time"
)

const NumberOfFrames = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
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
