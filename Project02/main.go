package main

import (
	"machine"
	"time"
)

func main() {
	pin3 := machine.D3
	pin3.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pin4 := machine.D4
	pin4.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pin5 := machine.D5
	pin5.Configure(machine.PinConfig{Mode: machine.PinOutput})

	pin2 := machine.D2
	pin2.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		switchStateIsOn := pin2.Get()

		if !switchStateIsOn {
			pin3.Set(true)
			pin4.Set(false)
			pin5.Set(false)
		} else {
			pin3.Set(false)
			pin4.Set(false)
			pin5.Set(true)

			time.Sleep(time.Second)

			pin4.Set(true)
			pin5.Set(false)

			time.Sleep(time.Second)
		}
	}
}
