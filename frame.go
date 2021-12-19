package main

type frame struct {
	previousFrame *frame
	knockDowns    []int
	points        int
}

func (f *frame) roll(standingPins int) int {
	if standingPins == 0 {
		f.knockDowns = append(f.knockDowns, -1)
		return 0
	} else {
		knockedDownPins := getKnockedDownPins(standingPins)
		f.knockDowns = append(f.knockDowns, knockedDownPins)
		f.points += knockedDownPins

		return knockedDownPins
	}

}

func (f *frame) addApplicableBonusPoints(sourceRollNr int, knockedDownPins int) {
	if f.hadSpare() && sourceRollNr == 1 || f.hadStrike() {
		f.points += knockedDownPins
	}

	if sourceRollNr == 1 && f.hadStrike() && f.previousFrame != nil && f.previousFrame.hadStrike() {
		f.previousFrame.points += knockedDownPins
	}
}

func (f *frame) hadSpare() bool {
	return len(f.knockDowns) > 1 && f.knockDowns[1] > 0 && f.knockDowns[0]+f.knockDowns[1] == totalPins
}

func (f *frame) hadStrike() bool {
	return f.knockDowns[0] == totalPins
}

func (f *frame) getTotalPoints() int {
	result := f.points
	if f.previousFrame != nil {
		result += f.previousFrame.getTotalPoints()
	}

	return result
}
