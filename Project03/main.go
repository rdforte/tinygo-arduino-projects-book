package main

import (
	"machine"
	"math"
	"time"
)

const (
	baseLineTemp = 20.0
	increment    = 2.0
	tempLow      = baseLineTemp + increment
	tempMid      = tempLow + increment
	tempHigh     = tempMid + increment
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
		// ADC is read in using 16-bit resolution, so the value will be between 0 and 65535.
		// This will be a fraction of the input voltage (0-5V), depending on the voltage divider.
		// To get the voltage, we divide the ADC value by the maximum value of unsigned 16-bit ADC (65535) and multiply by 5V
		volts := 5.0
		sensorVoltage := float64(A0.Get()) / float64(math.MaxUint16) * volts
		temperatureCelcius := (sensorVoltage - 0.5) * 100
		println(int(temperatureCelcius)) // truncate degree

		if temperatureCelcius >= tempLow && temperatureCelcius < tempMid {
			pin2.High()
			pin3.Low()
			pin4.Low()
		} else if temperatureCelcius >= tempMid && temperatureCelcius < tempHigh {
			pin2.High()
			pin3.High()
			pin4.Low()
		} else if temperatureCelcius >= tempHigh {
			pin2.High()
			pin3.High()
			pin4.High()
		}

		time.Sleep(time.Second)
	}
}
