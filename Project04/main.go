package main

import (
	"machine"
	"time"
)

func main() {
	greenLEDPin := machine.D9
	greenLEDPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	greenLEDPin.Low()

	redLEDPin := machine.D10
	redLEDPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	redLEDPin.Low()

	blueLEDPin := machine.D11
	blueLEDPin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	blueLEDPin.Low()

	machine.InitADC()
	redSensorPin := machine.ADC{Pin: machine.ADC0}
	greenSensorPin := machine.ADC{Pin: machine.ADC1}
	blueSensorPin := machine.ADC{Pin: machine.ADC2}

	for {
		redSensorValue := redSensorPin.Get() // 0-65535
		time.Sleep(time.Millisecond * 5)
		greeSensorValue := greenSensorPin.Get()
		time.Sleep(time.Millisecond * 5)
		blueSensorValue := blueSensorPin.Get()

		println("---- Raw sensor values ----")
		println("Red: %d", redSensorValue)
		println("Green: %d", greeSensorValue)
		println("Blue: %d", blueSensorValue)

		time.Sleep(time.Second * 5)
	}

	// for {
	// 	// ADC is read in using 16-bit resolution, so the value will be between 0 and 65535.
	// 	// This will be a fraction of the input voltage (0-5V), depending on the voltage divider.
	// 	// To get the voltage, we divide the ADC value by the maximum value of unsigned 16-bit ADC (65535) and multiply by 5V
	// 	volts := 5.0
	// 	sensorVoltage := float64(A0.Get()) / float64(math.MaxUint16) * volts
	// 	temperatureCelcius := (sensorVoltage - 0.5) * 100
	// 	println(int(temperatureCelcius)) // truncate degree
	//
	// 	if temperatureCelcius >= tempLow && temperatureCelcius < tempMid {
	// 		pin2.High()
	// 		redLEDPin.Low()
	// 		blueLEDPin.Low()
	// 	} else if temperatureCelcius >= tempMid && temperatureCelcius < tempHigh {
	// 		pin2.High()
	// 		redLEDPin.High()
	// 		blueLEDPin.Low()
	// 	} else if temperatureCelcius >= tempHigh {
	// 		pin2.High()
	// 		redLEDPin.High()
	// 		blueLEDPin.High()
	// 	}
	//
	// 	time.Sleep(time.Second)
	// }
}
