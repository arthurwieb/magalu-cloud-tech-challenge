package aggregator

import (
	"fmt"

	"github.com/arthurwieb/magalu-cloud-tech-challenge/pulsegenerator"
)

type Aggregator struct {
	pulseChan chan pulsegenerator.Pulse
	data      map[string]float64 // agrega dados internamente
}

func NewAggregator() *Aggregator {
	agg := &Aggregator{
		pulseChan: make(chan pulsegenerator.Pulse),
		data:      make(map[string]float64), // criar estrutura json
	}

	go agg.start()
	return agg
}

func (a *Aggregator) AddPulse(p pulsegenerator.Pulse) {
	a.pulseChan <- p
}

func (a *Aggregator) start() {
	// for {
	// 	// select seleciona alg
	// 	select {
	// 	case p := <-a.pulseChan:
	// 		key := p.Tenant + "|" + p.ProductSKU + "|" + p.UsageUnit
	// 		a.data[key] += p.UsedAmount
	// 	default:
	// 		close(a.pulseChan)
	// 	}
	// }

	//TODO IMPLEMENTAR OS DIAS
	for p := range a.pulseChan {
		fmt.Println("Pulso gerado: ", p)
		key := p.Tenant + "|" + p.ProductSKU + "|" + p.UsageUnit
		a.data[key] += p.UsedAmount
	}
}

func (a *Aggregator) Stop() {
	close(a.pulseChan)
}

func (a *Aggregator) GetTotals() map[string]float64 {
	a.Stop()
	return a.data
}
