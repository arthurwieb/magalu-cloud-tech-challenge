package main

import (
	"fmt"

	"github.com/arthurwieb/magalu-cloud-tech-challenge/aggregator"
	"github.com/arthurwieb/magalu-cloud-tech-challenge/pulsegenerator"
)

func main() {

	agg := aggregator.NewAggregator()

	// for que simula os dias no mes
	for i := 1; i <= 30; i++ {
		// simular pulsos diários
		for range 20 {
			randomPulse := pulsegenerator.GenerateRandomPulse(i)
			agg.AddPulse(randomPulse)
		}

	}

	fmt.Println(agg.GetTotals())

}
