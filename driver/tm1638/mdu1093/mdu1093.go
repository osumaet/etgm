// Package mdu1093 implements driver for TM1638 based board
package mdu1093

import (
	"github.com/osumaet/etgm/driver/tm1638"
	"machine"
)

const ()

// Device is representation of MDU1093 board
type Device struct {
	// led is bitmap of LEDs
	led byte
	// displayMemory contains one byte per 7-segment display
	displayMemory [8]byte
	// brightness hold last specified brightness level
	brightness byte
	// keyBuffer store 4 bytes of keypad scan data
	keyBuffer [4]byte
	// ic is reference on
	ic tm1638.Device
}

// New creates new MDU1093 device
func New(strobe machine.Pin, clock machine.Pin, data machine.Pin) Device {
	return Device{ic: tm1638.New(strobe, clock, data), brightness: tm1638.MaxBrightness}
}

// Configure prepare board to operations
func (d *Device) Configure() {
	d.ic.Configure()
	d.ic.SetDisplayBrightness(d.brightness)
}
