package main

import "time"

type Pulse struct {
	Tenant    string
	Service   string
	Amount    float64
	Unit      string
	Timestamp time.Time
}

// estrutura do pulse conforme contexto
// type Pulse struct {
// 	Tenant     string
// 	ProductSKU string
// 	UsedAmout  float64
// 	UsageUnit  string
// 	Timestamp  time.Time
// }

// poderemos utilizar esse pulse para storage ou network_egress (download) exemplos a baixo.

// storagePulse := Pulse{
//     Tenant:    "tenant_xpto",
//     ProductSKU: "storage_1gb",
//     UsedAmount: 60,
//     UsageUnit:  "GB x seg",
// }

// downloadPulse := Pulse{
//     Tenant:    "tenant_xpto",
//     ProductSKU: "network_egress",
//     UsedAmount: 307,
//     UsageUnit:  "KB",
// }
