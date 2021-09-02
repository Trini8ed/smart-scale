package main

import (
	"fmt"
	"log"
	"time"

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

func main() {
	/*
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

		var data int
		for i := 0; i < 3; i++ {
			time.Sleep(200 * time.Microsecond)

			data, err := hx711.ReadDataRaw()
			if err != nil {
				fmt.Println("ReadDataRaw error:", err)
				continue
			}

			fmt.Println(data)
		}
	*/
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
	//fmt.Printf("%v\n", read[1:])

	time.Sleep(time.Microsecond * 100)
	/*******************************************************************/
	clearScreen := []byte{0xFE, 0x62}
	read2 := make([]byte, len(clearScreen))
	if err != nil {
		fmt.Println("cannot open LCD device", err)
		return
	}

	if err := c.Tx(clearScreen, read2); err != nil {
		fmt.Println("LCD Turned On!", err)
		log.Fatal(err)
	}

	/*******************************************************************/
	//display on screen
	/*
		var data int = 01234567
		stringNumber := strconv.Itoa(data)
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

			//Move cursor right one place
			moveCursor := []byte{0xFE, 0x4A}
			read4 := make([]byte, len(moveCursor))
			if err != nil {
				fmt.Println("cannot display LCD device", err)
				return
			}
			if err := c.Tx(moveCursor, read4); err != nil {
				log.Fatal(err)
				fmt.Println("Move Cursor!", err)
				time.Sleep(time.Microsecond * 100)
			}
		}
	*/
}
