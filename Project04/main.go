package main

import (
	"time"

	"machine"
)

func main() {
	pwmGreen := machine.Timer1
	pwmGreen.Configure(machine.PWMConfig{})

	pwmRed := machine.Timer1
	pwmRed.Configure(machine.PWMConfig{})

	pwmBlue := machine.Timer2
	pwmBlue.Configure(machine.PWMConfig{})

	greenLEDPin := machine.D9
	greenChan, err := pwmGreen.Channel(greenLEDPin)
	if err != nil {
		println("err greenChan:", err)
		return
	}

	redLEDPin := machine.D10
	redChan, err := pwmRed.Channel(redLEDPin)
	if err != nil {
		println("err redChan:", err)
		return
	}

	blueLEDPin := machine.D11
	blueChan, err := pwmBlue.Channel(blueLEDPin)
	if err != nil {
		println("err blueChan:", err)
		return
	}

	machine.InitADC()
	redSensorPin := machine.ADC{Pin: machine.ADC0}
	greenSensorPin := machine.ADC{Pin: machine.ADC1}
	blueSensorPin := machine.ADC{Pin: machine.ADC2}

	for {
		redSensorValue := float32(redSensorPin.Get())
		greeSensorValue := float32(greenSensorPin.Get())
		blueSensorValue := float32(blueSensorPin.Get())

		total := redSensorValue + greeSensorValue + blueSensorValue
		redPercent := redSensorValue / total
		greenPercent := greeSensorValue / total
		bluePercent := blueSensorValue / total

		redOutput := redPercent * float32(pwmGreen.Top())
		greenOutput := greenPercent * float32(pwmGreen.Top())
		blueOutput := bluePercent * float32(pwmBlue.Top())

		pwmGreen.Set(greenChan, uint32(greenOutput))
		pwmRed.Set(redChan, uint32(redOutput))
		pwmBlue.Set(blueChan, uint32(blueOutput))
		time.Sleep(time.Second)
	}
}
