package bowling

import "errors"

const RollsPerFrame = 2

type Frame struct {
	PreviousFrame *Frame
	FrameNr       int
	KnockDowns    []int
	Points        int
}

func (f *Frame) ExecuteRolls() error {
	if f.KnockDowns != nil {
		return errors.New("the frame was already played")
	}

	for currRollNumber := 0; currRollNumber < RollsPerFrame; currRollNumber++ {
		knockedDownPins := 7
		f.KnockDowns = append(f.KnockDowns, knockedDownPins)
		f.Points += knockedDownPins
	}

	return nil
}
