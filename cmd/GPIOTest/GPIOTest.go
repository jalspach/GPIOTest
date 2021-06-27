package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"

	"github.com/yryz/ds18b20"
)

func sweep() int {
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
}

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

func tempprint() {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			fmt.Printf("sensor: %s temperature: %.2f°C\n", sensor, t)
		}
	}
}
func tempc() float64 {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)

	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			return t
		}
	}
	return 0
}
func main() {

	sweep()
	error()
	tempc()
	tempf := (tempc() * (9 / 5)) + 32
	fmt.Printf("The temperature is: %.2f°C\n", tempc())
	fmt.Printf("The temperature is: %.2f°F\n", tempf)
	tempprint()
}
