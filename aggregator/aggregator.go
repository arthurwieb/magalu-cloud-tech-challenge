package aggregator

import (
	"strconv"
	"sync"

	"github.com/arthurwieb/magalu-cloud-tech-challenge/pulsegenerator"
)

type Aggregator struct {
	pulseChan chan pulsegenerator.Pulse
	data      map[string]float64 // agrega dados internamente
	mu        sync.Mutex         // para resolver problema de concorrencia
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
	for p := range a.pulseChan {
		// fmt.Println("Pulso gerado: ", p)
		key := p.Tenant + "|" + p.ProductSKU + "|" + p.UsageUnit + "| dia:" + strconv.Itoa(p.Day) + "|"
		a.mu.Lock()
		a.data[key] += p.UsedAmount
		a.mu.Unlock()
	}
}

func (a *Aggregator) Stop() {
	close(a.pulseChan)
}

func (a *Aggregator) GetTotals() map[string]float64 {
	a.Stop()
	a.mu.Lock()
	defer a.mu.Unlock()

	// cria uma copia para ler a data do aggregator, não fazendo isso pode ocorrer erro de concorrencia
	result := make(map[string]float64)
	for k, v := range a.data {
		result[k] = v
	}
	return result
}
