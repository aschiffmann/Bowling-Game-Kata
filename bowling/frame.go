package bowling

import (
	"math/rand"
)

const RollsPerFrame = 2
const NumberOfPins = 10

type Frame struct {
	PreviousFrame *Frame
	FrameNr       int
	KnockDowns    []int
	Points        int
}

func (f *Frame) ExecuteRolls() {
	for rollNr := 0; rollNr < RollsPerFrame; rollNr++ {
		knockedDownPins := doTheRoll(NumberOfPins - f.Points)
		f.KnockDowns = append(f.KnockDowns, knockedDownPins)
		f.Points += knockedDownPins
	}
}

func doTheRoll(standingPins int) int {
	return rand.Intn(standingPins + 1)
}
