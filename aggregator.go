package main

import "fmt"

type Aggregator struct {
	data map[string]float64
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		data: make(map[string]float64),
	}
}

// adiciona um pulso a um agregador
func (a *Aggregator) addPulse(p Pulse) {
	key := p.Tenant + "-" + p.Service
	a.data[key] += p.Amount
}

func (a *Aggregator) printTotal() {
	fmt.Println("total")
	for key, total := range a.data {
		fmt.Printf("- %s: %.2f\n", key, total)

	}
}
