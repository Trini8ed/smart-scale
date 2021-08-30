package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MichaelS11/go-hx711"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

func main() {

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

	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Microsecond)

		data, err := hx711.ReadDataRaw()
		if err != nil {
			fmt.Println("ReadDataRaw error:", err)
			continue
		}

		fmt.Println(data)
	}

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

	// Write 0x10 to the device, and read a byte right after.

	// turns on the display
	displayOn := []byte{0x41, 0xFE}
	read := make([]byte, len(displayOn))
	if err != nil {
		fmt.Println("cannot open LCD device", err)
		return
	}

	if err := c.Tx(displayOn, read); err != nil {
		log.Fatal(err)
	}
	/*
		if _, err := NewRecordRaw(data); err != nil {
			fmt.Println("cannot read data", err)
			return
		}
	*/
	// Use read.
	fmt.Printf("%v\n", read[1:])
}
