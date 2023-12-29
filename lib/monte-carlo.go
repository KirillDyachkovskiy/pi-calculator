package lib

import (
	"math"
	"math/rand"
	"pi-calculator/utils"
	"sync"
)

func calculateThread(iterations uint, group *sync.WaitGroup, circledCount chan uint) {
	defer group.Done()

	circledLocalCount := uint(0)

	for i := uint(0); i < iterations; i++ {
		x := rand.Float64()
		y := rand.Float64()

		isCircled := math.Pow(x, 2)+math.Pow(y, 2) <= 1

		if isCircled {
			circledLocalCount++
		}
	}

	circledCount <- circledLocalCount
}

func CalculatePi(threads uint, iterations uint) float64 {
	defer utils.Timer("Monte-Carlo")()

	circledCount := make(chan uint)
	group := new(sync.WaitGroup)

	group.Add(int(threads))

	localIterations := iterations / threads

	for i := uint(0); i < threads; i++ {
		go calculateThread(localIterations, group, circledCount)
	}

	go func() {
		group.Wait()
		close(circledCount)
	}()

	totalCircledCount := uint(0)
	for count := range circledCount {
		totalCircledCount += count
	}

	return 4 * float64(totalCircledCount) / float64(iterations)
}
