package aggregator

import (
	"encoding/json"
	"sync"

	"github.com/arthurwieb/magalu-cloud-tech-challenge/pulsegenerator"
)

type Aggregator struct {
	pulseChan chan pulsegenerator.Pulse
	data      map[string]*AggregatedPulse
	mu        sync.Mutex // para resolver problema de concorrencia
}

type AggregatedPulse struct {
	Tenant     string          `json:"tenant"`
	ProductSKU string          `json:"product_sku"`
	UsageUnit  string          `json:"usage_unit"`
	DailyUsage map[int]float32 `json:"daily_usage"`
}

func NewAggregator() *Aggregator {
	// verificar sobre buffer nesses
	agg := &Aggregator{
		pulseChan: make(chan pulsegenerator.Pulse),
		data:      make(map[string]*AggregatedPulse), // criar estrutura json
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
		key := p.Tenant + "|" + p.ProductSKU + "|" + p.UsageUnit
		a.mu.Lock()

		// se essa chave ainda não existir, cria ela em nosso objeto
		if a.data[key] == nil {
			a.data[key] = &AggregatedPulse{
				Tenant:     p.Tenant,
				ProductSKU: p.ProductSKU,
				UsageUnit:  p.UsageUnit,
				DailyUsage: make(map[int]float32),
			}
		}
		a.data[key].DailyUsage[p.Day] += p.UsedAmount

		a.mu.Unlock()
	}
}

func (a *Aggregator) Stop() {
	close(a.pulseChan)
}

func (a *Aggregator) GetTotals() ([]byte, error) {
	a.Stop()
	a.mu.Lock()
	defer a.mu.Unlock()

	// cria uma copia(slice) para ler a data do aggregator, não fazendo isso pode ocorrer erro de concorrencia
	var results []*AggregatedPulse
	for _, v := range a.data {
		results = append(results, v)
	}

	// JSON serialize
	jsonData, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
