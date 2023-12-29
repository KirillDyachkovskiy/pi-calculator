package main

import (
	"flag"
	"fmt"
	"os"
	"pi-calculator/lib"
)

type options struct {
	threads    uint
	iterations uint
}

func main() {
	options := getOptions()

	fmt.Printf("threads %v\n", options.threads)
	fmt.Printf("iterations  %v\n", options.iterations)

	result := lib.CalculatePi(options.threads, options.iterations)

	fmt.Printf("approximate number of PI %v\n", result)
}

func getOptions() options {
	iterationsFlag := flag.Uint("iterations", 10_000, "The number of points - the accuracy of calculations depends on it")
	threadsFlag := flag.Uint("threads", 8, "The number of threads")
	flag.Parse()

	iterations := *iterationsFlag
	threads := *threadsFlag

	if threads == uint(0) {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return options{threads, iterations}
}
