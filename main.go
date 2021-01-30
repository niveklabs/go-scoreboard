package main

import (
	"machine"
	"time"
)

func main() {

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	button := machine.D7
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	debounce := false

	for {
		pushed := button.Get()

		if pushed {
			if debounce {
				continue
			}
			led.Low()
			println("BUTTON PUSHED")
			debounce = true
		} else {
			led.High()
			debounce = false
		}

		time.Sleep(time.Millisecond * 100)
	}
}
