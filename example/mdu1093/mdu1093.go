package main

import (
	"github.com/osumaet/etgm/driver/tm1638"
	"machine"
	"time"
)

func main() {
	tm := tm1638.NewDevice(machine.D7, machine.D9, machine.D8)
	tm.Configure()
	tm.SetDisplayBrightness(tm1638.MaxBrightness)

	for {
		for i := uint8(0); i < tm1638.MaxAddress; i++ {
			if i > 0 {
				tm.Open()
				tm.SendByte(tm1638.CmdFixedAddress)
				tm.Close()

				tm.Open()
				tm.SendByte(tm1638.CmdSetAddress | (i - 1))
				tm.SendByte(0x00)
				tm.Close()
			}
			tm.Open()
			tm.SendByte(tm1638.CmdFixedAddress)
			tm.Close()

			tm.Open()
			tm.SendByte(tm1638.CmdSetAddress | i)
			tm.SendByte(0x7F)
			tm.Close()

			time.Sleep(time.Millisecond * 250)
		}
		time.Sleep(time.Second)
		tm.ClearDisplayMemory()
	}
}
