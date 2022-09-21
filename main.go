package main

import (
	"fmt"
	"math"
	"time"
)

type ReviewItem struct {
	reps            int
	easeScore       float64
	daysBetweenReps float64
}

func NewReviewItem() *ReviewItem {
	ri := new(ReviewItem)
	ri.reps = 0
	ri.easeScore = 1
	ri.daysBetweenReps = 2.5
	return ri
}

func (ri *ReviewItem) NextPractice(quality int) time.Time {
	// TODO assert 0 <= quality <= 5
	ri.easeScore = math.Max(1.3, ri.easeScore+0.1-(5.0-float64(quality))*(0.08+(5.0-float64(quality))*0.02))

	if quality < 3 {
		ri.reps = 0
	} else {
		ri.reps += 1
	}

	if ri.reps <= 1 {
		ri.daysBetweenReps = 1
	} else if ri.reps == 2 {
		ri.daysBetweenReps = 5
	} else {
		ri.daysBetweenReps = math.Round(ri.daysBetweenReps * ri.easeScore)
	}

	daysToAdd, _ := time.ParseDuration(
		fmt.Sprintf("%.2fh", ri.daysBetweenReps*24),
	)

	return time.Now().Add(daysToAdd)
}

func main() {
	fmt.Println("We are up an running!")
	var rn *ReviewItem = NewReviewItem()
	fmt.Println(rn.NextPractice(4))
}
