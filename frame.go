package main

type frame struct {
	previousFrame *frame
	knockDowns    []int
	points        int
}

func (f *frame) roll(standingPins int) int {
	knockedDownPins := getKnockedDownPins(standingPins)
	f.knockDowns = append(f.knockDowns, knockedDownPins)
	f.points += knockedDownPins

	return knockedDownPins
}

func (f *frame) addApplicableBonusPoints(knockDownsRoll1 int, knockDownsRoll2 int) {
	if prev := f.previousFrame; prev != nil {
		if prev.hadSpare() {
			prev.points += knockDownsRoll1
		}

		if prev.hadStrike() {
			prev.points += knockDownsRoll1 + knockDownsRoll2

			if prev.previousFrame != nil && prev.previousFrame.hadStrike() {
				prev.previousFrame.points += knockDownsRoll1
			}
		}
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
