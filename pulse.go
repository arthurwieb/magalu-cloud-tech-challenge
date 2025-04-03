package main

import "time"

type Pulse struct {
	Tenant    string
	Service   string
	Amount    float64
	Unit      string
	Timestamp time.Time
}
