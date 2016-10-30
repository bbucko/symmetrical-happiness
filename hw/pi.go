package hw

import (
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/kidoman/embd/controller/hd44780"
	"fmt"
	"time"
)

type Pi struct {
	lcdPins    []embd.DigitalPin
	lightsPins []embd.DigitalPin
	lcd        *hd44780.HD44780
}

func Init() (pi *Pi) {
	pi = new(Pi)
	pi.lcdPins = make([]embd.DigitalPin, 0, 7)
	pi.lightsPins = make([]embd.DigitalPin, 0, 3)

	host, rev, err := embd.DetectHost()
	if err != nil {
		panic(err)
	}

	fmt.Println("Init ", host, " ", rev)

	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}

	pin1, err := embd.NewDigitalPin("P1_12")
	if (err != nil) {
		panic(err)
	}
	pin1.SetDirection(embd.Out)
	pin1.Write(embd.High)
	pi.lightsPins = append(pi.lightsPins, pin1)

	pin2, err := embd.NewDigitalPin("P1_16")
	if (err != nil) {
		panic(err)
	}
	pin2.SetDirection(embd.Out)
	pin2.Write(embd.High)
	pi.lightsPins = append(pi.lightsPins, pin2)

	pin3, err := embd.NewDigitalPin("P1_18")
	if (err != nil) {
		panic(err)
	}
	pin3.SetDirection(embd.Out)
	pin3.Write(embd.High)
	pi.lightsPins = append(pi.lightsPins, pin3)

	rs, err := embd.NewDigitalPin("P1_26")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, rs)

	en, err := embd.NewDigitalPin("P1_24")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, en)

	d4, err := embd.NewDigitalPin("P1_22")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, d4)

	d5, err := embd.NewDigitalPin("P1_23")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, d5)

	d6, err := embd.NewDigitalPin("P1_21")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, d6)

	d7, err := embd.NewDigitalPin("P1_19")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, d7)

	backlight, err := embd.NewDigitalPin("P1_11")
	if (err != nil) {
		panic(err)
	}
	pi.lcdPins = append(pi.lcdPins, backlight)

	pi.lcd, err = hd44780.NewGPIO(rs, en, d4, d5, d6, d7, backlight,
		hd44780.Negative,
		hd44780.RowAddress20Col,
		hd44780.TwoLine)
	if err != nil {
		panic(err)
	}

	pi.lcd.Clear()
	pi.Write("Hello Dave", 0)

	return
}

func (pi *Pi) Close() {
	for i, pin := range pi.lightsPins {
		println("Removing pin#", i, "for PIN#", pin.N())
		pin.Close()
	}

	pi.lcd.Close()
	embd.CloseGPIO()
}

func (pi *Pi) Write(msg string, lineNo int) {
	pi.lcd.SetCursor(0, lineNo)
	for _, b := range []byte(msg) {
		pi.lcd.WriteChar(b)
	}
}

func (pi *Pi) Blink(pinNo int) {
	pi.lightsPins[pinNo].Write(embd.High)
	time.Sleep(1 * time.Second)
	pi.lightsPins[pinNo].Write(embd.Low)
}