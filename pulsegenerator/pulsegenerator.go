package pulsegenerator

import (
	"math/rand"
)

var TENANT_ID = []string{"1", "2", "3", "4", "5"}

// type Emitter interface {
// 	generateRandomPulse() Pulse
// }

type Pulse struct {
	Tenant     string
	ProductSKU string
	UsedAmount float32
	UsageUnit  string
	Day        int // em um cenário real, seria um timestamp
}

// mock de dados de um pulse gerado
func GenerateRandomPulse(day int) Pulse {
	tenantName := "tenant_" + TENANT_ID[rand.Intn(len(TENANT_ID))]
	amount := rand.Float32()*100 + 1 // ajustar isso, parece super não efetivo

	var newPulse Pulse
	if rand.Intn(2) > 0 {
		newPulse = generateStoragePulse(tenantName, amount, day)
	} else {
		newPulse = generateNetworkPulse(tenantName, amount, day)
	}

	return newPulse
}

func generateStoragePulse(Tenant string, amount float32, day int) Pulse {

	return Pulse{
		Tenant:     Tenant,
		ProductSKU: "storage_1gb",
		UsedAmount: amount * 60,
		UsageUnit:  "GB x seg",
		Day:        day,
	}
}

func generateNetworkPulse(Tenant string, amount float32, day int) Pulse {

	return Pulse{
		Tenant:     Tenant,
		ProductSKU: "network_egress",
		UsedAmount: amount,
		UsageUnit:  "MB",
		Day:        day,
	}
}
