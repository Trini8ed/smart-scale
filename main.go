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

func main() {
	//calibrate()
	//var data [5]float64
	//getWeight(data)
	/*
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
		/*
			// SetGain default is 128
			// Gain of 128 or 64 is input channel A, gain of 32 is input channel B
			// hx711.SetGain(128)

			// make sure to use your values from calibration above
			hx711.AdjustZero = 5507
			hx711.AdjustScale = 30904
			/*
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
				for i := 0; i < 5; i++ {
					fmt.Println(int(data[i]))
				}
	*/

	//Digi Scale
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

	defer hx711.Shutdown()

	err = hx711.Reset()
	if err != nil {
		fmt.Println("Reset error:", err)
		return
	}
	hx711.AdjustZero = 5507
	hx711.AdjustScale = 30904
	var data [4]int
	for i := 0; i < 4; i++ {
		time.Sleep(200 * time.Microsecond)

		data[i], err = hx711.ReadDataRaw()
		if err != nil {
			fmt.Println("ReadDataRaw error:", err)
			continue
		}

		fmt.Println(data[i])
	}

	/*******************************************************************/
	//SPI
	// Make sure periph is initialized.

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

	/*******************************************************************/
	//LCD Screen activate
	// turns on the display

	displayOn := []byte{0xFE, 0x41}
	read := make([]byte, len(displayOn))
	if err != nil {
		fmt.Println("cannot open LCD device", err)
		return
	}

	if err := c.Tx(displayOn, read); err != nil {
		fmt.Println("LCD Turned On!", err)
		log.Fatal(err)
	}

	// Use read.

	time.Sleep(time.Microsecond * 100)
	/*******************************************************************/

	clearScreen := []byte{0xFE, 0x51}
	read2 := make([]byte, len(clearScreen))
	if err != nil {
		fmt.Println("cannot open LCD device", err)
		return
	}

	if err := c.Tx(clearScreen, read2); err != nil {
		fmt.Println("LCD Turned On!", err)
		log.Fatal(err)
	}
	time.Sleep(time.Microsecond * 100)

	/*******************************************************************/
	//display on screen
	//stringNumber := floattostr(data[4])
	stringNumber := strconv.Itoa(int(data[3]))
	runedNumbers := []rune(stringNumber)

	for _, r := range runedNumbers {
		fmt.Printf("Rune: %v Hex: 0x%x\n", strconv.QuoteRune(r), characterMap[r])

		// display number on LCD screen
		displaynumber := []byte{characterMap[r]}
		read3 := make([]byte, len(displaynumber))
		if err != nil {
			fmt.Println("cannot display LCD device", err)
			return
		}

		if err := c.Tx(displaynumber, read3); err != nil {
			log.Fatal(err)
			fmt.Println("display number!", err)
			return
		}
		time.Sleep(time.Microsecond * 100)
	}

	//var data int = 400
	database(int(data[3]))

}
