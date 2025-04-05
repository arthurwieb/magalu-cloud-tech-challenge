package main

import (
	"fmt"

	"github.com/arthurwieb/magalu-cloud-tech-challenge/aggregator"
	"github.com/arthurwieb/magalu-cloud-tech-challenge/pulsegenerator"
)

func main() {

	agg := aggregator.NewAggregator()

	for range 5 {
		randomPulse := pulsegenerator.GenerateRandomPulse(1)
		agg.AddPulse(randomPulse)
	}

	fmt.Println(agg.GetTotals())

}
