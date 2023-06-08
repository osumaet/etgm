/*
TM1638 is a chip manufactured by Titan Microelectronics.
It integrates MCU digital interface, data latch, LED drive, and keypad scanning circuit.
*/
package tm1638

import (
	"machine"
	"time"
)

const (
	/* Address increasing mode: automatic address increased */
	cmdAddressAutoIncrement = 0x40
	/* Address increasing mode: fixed address */
	CmdFixedAddress = 0x44
	/* Read key scan data */
	cmdReadKeyScan = 0x42
	/* Display off */
	cmdZeroBrightness = 0x80
	/* Display on command. Bits 0-2 may contain brightness value */
	cmdSetBrightness = 0x88
	/* Address Setting Command is used to set the address of the display memory. Bits 0-3 used for address value */
	CmdSetAddress = 0xC0
	/* Max valid address */
	MaxAddress = 0x0F
	/* MaxBrightness is max brightness value */
	MaxBrightness = 0x07
)

// Device Device wraps the pins of the TM1638.
type Device struct {
	/* STB - When this pin is "LO" chip accepting transmission */
	strobe machine.Pin
	/* CLK - DIO pin reads data at the rising edge and outputs at the falling edge */
	clock machine.Pin
	/* DIO - This pin outputs/inputs serial data */
	data machine.Pin
}

// NewDevice Create new TM1638 device
func NewDevice(strobe machine.Pin, clock machine.Pin, data machine.Pin) Device {
	return Device{strobe: strobe, clock: clock, data: data}
}

func (d *Device) Configure() {
	d.strobe.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.clock.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.data.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

// Open connection to IC
func (d *Device) Open() {
	d.strobe.Low()
	d.transmissionDelay()
}

// Close connection to IC
func (d *Device) Close() {
	d.strobe.High()
	d.transmissionDelay()
}

func (d *Device) SendByte(value uint8) {
	for i := 0; i < 8; i++ {
		d.clock.Low()
		d.transmissionDelay()
		d.data.Set((value & (1 << i)) > 0)
		d.transmissionDelay()
		d.clock.High()
		d.transmissionDelay()
	}
}

/*
Keyboard scan matrix has 3 rows (K1-K3) and 8 (KS1-KS8) columns.
Each button connected to one K and KS line of IC.

|---------|---------------------|---------------------|
|  Mapping table                                      |
|---------|---------------------|---------------------|
|Columns: |- K3 - K2 - K1 - X  -|- K3 - K2 - K1 - X  -|
|---------|---------------------|---------------------|
|Bit:     |- B0 - B1 - B2 - B3 -|- B4 - B5 - B6 - B7 -|
|---------|---------------------|---------------------|
|BYTE-1:  |         KS1         |         KS2         |
|BYTE-2:  |         KS3         |         KS4         |
|BYTE-3:  |         KS5         |         KS6         |
|BYTE-4:  |         KS7         |         KS8         |
|---------|---------------------|---------------------|
*/
func (d *Device) ReadKeyboard(buffer *[4]uint8) {
	d.Open()
	d.SendByte(cmdReadKeyScan)
	d.data.Configure(machine.PinConfig{Mode: machine.PinInput})
	d.transmissionDelay()
	var element uint8 = 0
	for index := range buffer {
		for bitIndex := 0; bitIndex < 8; bitIndex++ {
			d.clock.Low()
			d.transmissionDelay()
			if d.data.Get() {
				element |= 1 << bitIndex
			}
			d.transmissionDelay()
			d.clock.High()
			d.transmissionDelay()
		}
		buffer[index] = element
		element = 0
	}
	d.data.Configure(machine.PinConfig{Mode: machine.PinOutput})
	d.Close()
}

func (d *Device) SetDisplayBrightness(value uint8) {
	d.Open()
	if value == 0 {
		d.SendByte(cmdZeroBrightness)
	} else {
		if value > MaxBrightness {
			value = MaxBrightness
		}
		d.SendByte(cmdSetBrightness | value)
	}
	d.Close()
}

func (d *Device) ClearDisplayMemory() {
	d.Open()
	d.SendByte(cmdAddressAutoIncrement)
	d.Close()

	d.Open()
	d.SendByte(CmdSetAddress)
	for i := 0; i < 16; i++ {
		d.SendByte(0)
	}
	d.Close()
}

func (d *Device) transmissionDelay() {
	time.Sleep(time.Microsecond * time.Duration(5))
}
