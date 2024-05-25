package main

import (
	"machine"
	"math"
	"time"
)

const (
	sensorPin    float64 = 0
	baseLineTemp float64 = 20.0
)

func main() {
	pin2 := machine.D2
	pin2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin2.Low()

	pin3 := machine.D3
	pin3.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin3.Low()

	pin4 := machine.D4
	pin4.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin4.Low()

	machine.InitADC()
	A0 := machine.ADC{Pin: machine.ADC0}

	for {
		voltage := float64(A0.Get()) / float64(math.MaxUint16) * 5
		temperatureCelcius := (voltage - 0.5) * 100
		println(int(temperatureCelcius)) // truncate degree
		time.Sleep(time.Second)
	}
}
