package main

import (
	"math/rand"
	"time"
)

const totalFrames = 10
const totalPins = 10

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	game := playNewGame()
	printResults(game)
}

func playNewGame() []*frame {
	var game []*frame
	var previousFrame *frame

	for frameNr := 1; frameNr <= totalFrames; frameNr++ {
		currFrame := playNewFrame(frameNr, previousFrame)

		game = append(game, currFrame)
		previousFrame = currFrame
	}
	return game
}

func playNewFrame(frameNr int, previousFrame *frame) *frame {
	currFrame := frame{previousFrame: previousFrame}

	knockDowns1 := currFrame.roll(totalPins)
	knockDowns2 := currFrame.roll(totalPins - knockDowns1)

	if previousFrame != nil {
		previousFrame.addApplicableBonusPoints(1, knockDowns1)
		previousFrame.addApplicableBonusPoints(2, knockDowns2)
	}

	if (currFrame.hadStrike() || currFrame.hadSpare()) && frameNr == totalFrames {
		currFrame.roll(totalPins)
	}

	return &currFrame
}
