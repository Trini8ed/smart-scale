package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MichaelS11/go-hx711"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

var characterMap = map[rune]byte{
	'0': 0x30,
	'1': 0x31,
	'2': 0x32,
	'3': 0x33,
	'4': 0x34,
	'5': 0x35,
	'6': 0x36,
	'7': 0x37,
	'8': 0x38,
	'9': 0x39,
	'a': 0x16,
	'b': 0x26,
	'c': 0x36,
	'd': 0x46,
}

func calibrate() {
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}

	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711.SetGain(128)

	var weight1 float64
	var weight2 float64

	weight1 = 0
	weight2 = 202

	hx711.GetAdjustValues(weight1, weight2)
}

func floattostr(input_num float64) string {

	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'g', 1, 64)
}

func lcdDisplay(data rune, characterMap map[rune]byte) {
	if _, err := host.Init(); err != nil {
		fmt.Println("host.Init error:", err)
		return
	}

	// Use spireg SPI port registry to find the first available SPI bus.
	p, err := spireg.Open("")
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer p.Close()

	// the spi.Port into a spi.Conn so it can be used for communication.
	c, err := p.Connect(physic.KiloHertz, spi.Mode3, 8)
	if err != nil {
		log.Fatal("Connect: ", err)
	}

	lcdDisplay := []byte{characterMap[data]}
	read := make([]byte, len(lcdDisplay))

	if err := c.Tx(lcdDisplay, read); err != nil {
		fmt.Println("LCD Turned On!", err)
		log.Fatal(err)
	}

}
func getWeight(data [5]float64) {
	err := hx711.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	hx711, err := hx711.NewHx711("GPIO6", "GPIO5")
	if err != nil {
		fmt.Println("NewHx711 error:", err)
		return
	}

	// SetGain default is 128
	// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
	// hx711.SetGain(128)

	// make sure to use your values from calibration above
	hx711.AdjustZero = 5507
	hx711.AdjustScale = 30904

	//var data[5] float64
	for i := 0; i < 5; i++ {

		time.Sleep(200 * time.Microsecond)

		data[i], err = hx711.ReadDataMedian(11)
		if err != nil {
			fmt.Println("ReadDataMedian error:", err)
			continue
		}
		fmt.Println(data[i])

	}
	return
}
