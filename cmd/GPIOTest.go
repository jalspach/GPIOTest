package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func sweep(speed int) int {
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
	for x := 0; x < speed; x++ {
		green.Toggle()
		time.Sleep(time.Second / 5)
		yellow.Toggle()
		time.Sleep(time.Second / 5)
		red.Toggle()
		time.Sleep(time.Second / 5)

	}
	return 0
}

// Set the status of the LED's based on the the Hex value. If an invalid number is passed will return 0X_BEEF if everything worked it will return 0x_FFFF
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
		yellow.Low()
		green.Low()
	case 0x_1:
		//Red solid
		red.High()
		yellow.Low()
		green.Low()
	case 0x_2:
		//Yellow solid 
		 red.Low()
		 yellow.High()
		 green.Low()
	case 0x_3:
		//Red and Yellow solid
		red.High()
		yellow.High()
		green.Low()
	case 0x_4:
		//Green solid
		red.Low()
		yellow.Low()
		green.High()
	
	case 0x_5:
		//Red and Green solid
		red.High()
		yellow.Low()
		green.High()
	case 0x_6:
		//Yellow and Green solid
		red.Low()
		yellow.High()
		green.High()
	case 0x_7:
		//Red, Green and Yellow solid
		red.High()
		yellow.High()
		green.High()
	case 0x_8:
		//Not Valid (all off but still flashing)
	case 0x_9:
		//Red flashing
		red.Toggle()
		yellow.Low()
		green.Low()
	case 0x_A:
		//Yellow flashing
		red.Low()
		yellow.Toggle()
		green.Low()
	
	case 0x_B:
		//Red and Yellow flashing
		red.Toggle()
		yellow.Toggle()
		green.Low()
	case 0x_C:
		//Green flashing
		red.Low()
		yellow.Low()
		green.Toggle()
	case 0x_D:
		//Red Green flashing
		red.Toggle()
		yellow.Low()
		green.Toggle()

	case 0x_E:
		//Yellow and Green flashing
		red.Low()
		yellow.Toggle()
		green.Toggle()

	case 0x_F:
		//Red, Yellow and Green flashing
		red.Toggle()
		yellow.Toggle()
		green.Toggle()
	default:
		return 0x_BEEF
	}
	return 0x_FFFF
}

func main() {

	sweep(20)


}
