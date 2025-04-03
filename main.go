package main

import "time"

func main() {
	agg := NewAggregator()

	//simulando pulsos que iriam vir de uma API/queue
	pulses := []Pulse{
		{
			Tenant:    "cliente-1",
			Service:   "download",
			Amount:    1824,
			Unit:      "KB",
			Timestamp: time.Now(),
		},
		{
			Tenant:    "cliente-2",
			Service:   "download",
			Amount:    307,
			Unit:      "KB",
			Timestamp: time.Now(),
		},
		{
			Tenant:    "cliente-3",
			Service:   "storage",
			Amount:    1024,
			Unit:      "GB x seg",
			Timestamp: time.Now(),
		},
		{
			Tenant:    "cliente-4",
			Service:   "storage",
			Amount:    2048,
			Unit:      "GB x seg",
			Timestamp: time.Now(),
		},
	}

	for _, pulse := range pulses {
		agg.addPulse(pulse)
	}

	agg.printTotal()
}
