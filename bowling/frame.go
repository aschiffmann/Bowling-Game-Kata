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

		if f.PreviousFrame != nil {
			f.PreviousFrame.addBonusPoints(rollNr+1, knockedDownPins)
		}
	}
}

func (f *Frame) addBonusPoints(rollNr int, knockedDownPins int) {
	if f.HadStrike() || f.HadSpare() && rollNr == 1 {
		f.Points += knockedDownPins
	}
}

func (f *Frame) HadSpare() bool {
	return f.KnockDowns[0] > 0 && f.KnockDowns[1] > 0 && f.KnockDowns[0]+f.KnockDowns[1] == NumberOfPins
}

func (f *Frame) HadStrike() bool {
	return f.KnockDowns[0] == NumberOfPins || f.KnockDowns[1] == NumberOfPins
}

func doTheRoll(standingPins int) int {
	return rand.Intn(standingPins + 1)
}
