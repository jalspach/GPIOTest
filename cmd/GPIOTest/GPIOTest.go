package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"

	"github.com/yryz/ds18b20"
)

/*func sweep() int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()

	yellow := rpio.Pin(27)
	yellow.Output()

	red := rpio.Pin(22)
	red.Output()

	red.Low()
	green.Low()
	yellow.Low()

	for x := 0; x < 20; x++ {
		green.Toggle()
		time.Sleep(time.Second / 5)
		yellow.Toggle()
		time.Sleep(time.Second / 5)
		red.Toggle()
		time.Sleep(time.Second / 5)

	}
	return 0
}*/

func error() int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()

	yellow := rpio.Pin(27)
	yellow.Output()

	red := rpio.Pin(22)
	red.Output()

	red.Low()
	green.Low()
	yellow.Low()

	for x := 0; x < 20; x++ {
		red.Toggle()
		time.Sleep(time.Second / 5)

	}

	return 0
}

func alert() int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()

	yellow := rpio.Pin(27)
	yellow.Output()

	red := rpio.Pin(22)
	red.Output()

	red.Low()
	green.Low()
	yellow.Low()

	for x := 0; x < 20; x++ {
		yellow.High()
		time.Sleep(time.Second / 5)

	}

	return 0
}

func good() int {
	fmt.Println("opening gpio")
	err := rpio.Open()
	if err != nil {
		panic(fmt.Sprint("unable to open gpio...must be root to run", err.Error()))
	}

	defer rpio.Close()

	green := rpio.Pin(17)
	green.Output()

	yellow := rpio.Pin(27)
	yellow.Output()

	red := rpio.Pin(22)
	red.Output()

	red.Low()
	green.Low()
	yellow.Low()

	for x := 0; x < 20; x++ {
		green.High()
		time.Sleep(time.Second / 5)

	}

	return 0
}

func tempc() float64 {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("sensor IDs: %v\n", sensors)

	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			return t
		}
	}
	return 0
}

func temp2led() int {
	var (
		curtemp float64 = tempc()
	)
	switch {
	case curtemp < 22.222:
		good()
	case curtemp > 29.444:
		error()
	default:
		alert()
	}
	return 0
}
func main() {

	/*sweep()
	error()

	tempprint()
	fmt.Printf("The temperature is: %.2f°C\n", tempc())
	fmt.Printf("The temperature is: %.2f°F\n", tempf)
	*/
	for i := 0; i < 100; i++ {

		tempf := fn0()
		temp2led()
		fmt.Printf("The temperature is: %.2f°C\n", tempc())
		fmt.Printf("The temperature is: %.2f°F\n", tempf)
		//sweep()
	}
}

func fn0() float64 {
	tempf := ((tempc() * (9 / 5)) + 32)
	return tempf
}
