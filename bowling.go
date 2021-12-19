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
	isLastFrame := frameNr == totalFrames
	currFrame := frame{previousFrame: previousFrame}

	knockDowns1 := currFrame.roll(totalPins)

	var knockDowns2 int
	if !currFrame.hadStrike() {
		knockDowns2 = currFrame.roll(totalPins - knockDowns1)
	} else if isLastFrame {
		knockDowns2 = currFrame.roll(totalPins)
	}

	currFrame.addApplicableBonusPoints(knockDowns1, knockDowns2)

	if (currFrame.hadStrike() || currFrame.hadSpare()) && isLastFrame {
		currFrame.roll(totalPins)
	}

	return &currFrame
}

var getKnockedDownPins = func(standingPins int) int {
	return rand.Intn(standingPins + 1)
}
