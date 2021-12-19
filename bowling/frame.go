package bowling

import (
	"bowling-game/config"
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
	knockedDownPins := doTheRoll(NumberOfPins - f.Points)
	f.KnockDowns = append(f.KnockDowns, knockedDownPins)
	f.Points += knockedDownPins

	if f.PreviousFrame != nil {
		f.PreviousFrame.addBonusPoints(1, knockedDownPins)
	}

	if !f.HadStrike() {
		knockedDownPins = doTheRoll(NumberOfPins - f.Points)
		f.KnockDowns = append(f.KnockDowns, knockedDownPins)
		f.Points += knockedDownPins

		if f.PreviousFrame != nil {
			f.PreviousFrame.addBonusPoints(2, knockedDownPins)
		}
	} else {
		f.KnockDowns = append(f.KnockDowns, -1)
	}

	if (f.HadStrike() || f.HadSpare()) && f.FrameNr == config.NumberOfFrames {
		knockedDownPins = doTheRoll(NumberOfPins)
		f.KnockDowns = append(f.KnockDowns, knockedDownPins)
		f.Points += knockedDownPins
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
	return f.KnockDowns[0] == NumberOfPins || len(f.KnockDowns) > 1 && f.KnockDowns[1] == NumberOfPins
}

func (f *Frame) GetTotalPoints() int {
	result := f.Points
	if f.PreviousFrame != nil {
		result += f.PreviousFrame.GetTotalPoints()
	}

	return result
}

func doTheRoll(standingPins int) int {
	return rand.Intn(standingPins + 1)
}
