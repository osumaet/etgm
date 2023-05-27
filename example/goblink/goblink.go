package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	firstChanel := make(chan uint8)
	go timerOne(firstChanel)
	for {
		select {
		case cnt := <-firstChanel:
			if cnt == 0 {
				led.High()
			} else {
				led.Low()
			}
		}
	}
}

func timerOne(ch chan<- uint8) {
	for {
		ch <- uint8(0)
		time.Sleep(500 * time.Millisecond)
		ch <- uint8(1)
		time.Sleep(500 * time.Millisecond)
	}
}
