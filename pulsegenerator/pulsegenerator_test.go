package pulsegenerator

import (
	"testing"
)

func TestNetworkPulseGenerator(t *testing.T) {

	p1 := generateNetworkPulse("p", 1, 1)
	p2 := Pulse{
		Tenant:     "p",
		ProductSKU: "network_egress",
		UsedAmount: 1,
		UsageUnit:  "MB",
		Day:        1,
	}

	if p1 != p2 {
		t.Errorf("Pulsos Network que eram para ser iguais estão diferentes, %v e %v", p1, p2)
	}
}

func TestStoragePulseGenerator(t *testing.T) {

	p1 := generateStoragePulse("p", 1, 1)
	p2 := Pulse{
		Tenant:     "p",
		ProductSKU: "storage_1gb",
		UsedAmount: 1 * 60,
		UsageUnit:  "GB x seg",
		Day:        1,
	}

	if p1 != p2 {
		t.Errorf("Pulsos Storage que eram para ser iguais estão diferentes, %v e %v", p1, p2)
	}
}
