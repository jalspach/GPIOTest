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
*/

func setstatus(hex int) int {
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

// Set the status of the leds
switch hex {

case 0x_0:
	//all off
	red.Low()
	green.Low()
	yellow.Low()
return 0x_0

case 0x_1:
	//Red solid
	red.High()
	green.Low()
	yellow.Low()
return 0x_1

case 0x_2:
	//Yellow solid
	red.Low()
	green.Low()
	yellow.High()
return 0x_2

case 0x_3:
	//Red and Yellow solid
	red.High()
	green.Low()
	yellow.High()
return 0x_3

case 0x_4:
	//Green solid
	red.Low()
	green.High()
	yellow.Low()
return 0x_4

case 0x_5:
	//Red and Green solid
	red.High()
	green.High()
	yellow.Low()
return 0x_5

case 0x_6:
	//Yellow and Green solid
	red.Low()
	green.High()
	yellow.High()
return 0x_6

case 0x_7:
	//Red, Green and Yellow solid
	red.High()
	green.High()
	yellow.High()
return 0x_6

case 0x_8:
	//Not Valid (all off but still flashing)
	return 0x_8

case 0x_9:
	//Red flashing
	red.High()
	green.Low()
	yellow.Low()
	for x := 0; x < 20; x++ {
		red.Toggle()
		time.Sleep(time.Second / 5)
	}
red.Low()
return 0x_9

case 0x_A:
	//Yellow flashing
	red.Low()
	green.Low()
	yellow.High()
	for x := 0; x < 20; x++ {
		yellow.Toggle()
		time.Sleep(time.Second / 5)
	}
	yellow.Low()
return 0x_A

case 0x_B:
	//Red and Yellow flashing
	red.High()
	green.Low()
	yellow.High()
	for x := 0; x < 20; x++ {
		red.Toggle()
		yellow.Toggle()
		time.Sleep(time.Second / 5)
	}
red.Low()
yellow.Low()
return 0x_B

case 0x_C:
	//Green flashing
	//Red and Yellow flashing
	red.Low()
	green.High()
	yellow.Low()
	for x := 0; x < 20; x++ {
		green.Toggle()
		time.Sleep(time.Second / 5)
	}
green.Low()
return 0X_C

case 0x_D:
	//Red Green flashing
	red.High()
	green.High()
	yellow.Low()
	for x := 0; x < 20; x++ {
		red.Toggle()
		green.Toggle()
		time.Sleep(time.Second / 5)
	}
red.Low()
green.Low()
return 0x_D

case 0x_E:
	//Yellow and Green flashing
	red.Low()
	green.High()
	yellow.High()
	for x := 0; x < 20; x++ {
		green.Toggle()
		yellow.Toggle()
		time.Sleep(time.Second / 5)
	}
green.Low()
yellow.Low()
return 0x_E

case 0x_F:
	//Red, Yellow and Green flashing
	red.High()
	green.High()
	yellow.High()
	for x := 0; x < 20; x++ {
		red.Toggle()
		green.Toggle()
		yellow.Toggle()
		time.Sleep(time.Second / 5)
	}
red.Low()
green.Low()
yellow.Low()
return 0x_F

default:
	return 0x_BEEF
}


func main()  {
	

	for i := 0; i < 0x_10; i++ {
		setstatus(i)
		time.Sleep (time.Second)
	}
}

/*
	for i := 0; i < 10; i++ {

		tempf := fn0()
		temp2led()
		fmt.Printf("The temperature is: %.2f°C\n", tempc())
		fmt.Printf("The temperature is: %.2f°F\n", tempf)
		//sweep()
	}
*/

	/*sweep()
	error()

	tempprint()
	fmt.Printf("The temperature is: %.2f°C\n", tempc())
	fmt.Printf("The temperature is: %.2f°F\n", tempf)
	*/

/*
func fn0() float64 {
	tempf := ((tempc() * (9 / 5)) + 32)
	return tempf
}
*/
