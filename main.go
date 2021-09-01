package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"time"

	"github.com/MichaelS11/go-hx711"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
)

/*
var numbers = map[int]byte{
	0: 0x03,
	1: 0x13,
	2: 0x23,
	3: 0x33,
	4: 0x43,
	5: 0x53,
	6: 0x63,
	7: 0x73,
	8: 0x83,
	9: 0x93,
}

var characters = map[string]byte{
	"a": 0x16,
	"b": 0x26,
	"c": 0x36,
	"d": 0x46,
}

func compare(data int, numbers var) {

	//convert int data into string
	sting_data := strconv.Itoa(data)
	for i := 0; i < 3; i++ {
		fmt.Println(sting_data[i : i+1])
		numbers[i]

	}

}
*/
func main() {

	debug.SetGCPercent(-1)
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
	//var data int
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

	// turns on the display
	displayOn := []byte{0xFE, 0x41}
	read := make([]byte, len(displayOn))
	if err != nil {
		fmt.Println("cannot open LCD device", err)
		return
	}

	if err := c.Tx(displayOn, read); err != nil {
		log.Fatal(err)
	}

	// Use read.
	fmt.Printf("%v\n", read[1:])
	time.Sleep(time.Microsecond * 100)

	// test print

	testing := []byte{0x13}
	read2 := make([]byte, len(testing))
	if err != nil {
		fmt.Println("cannot display", err)
		return
	}

	if err := c.Tx(testing, read2); err != nil {
		log.Fatal(err)
	}
}
