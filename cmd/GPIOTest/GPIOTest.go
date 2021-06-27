package main

import (
	"fmt"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func main() {
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

	for x := 0; x < 20; x++ {
		green.Toggle()
		time.Sleep(time.Second / 5)
		yellow.Toggle()
		time.Sleep(time.Second / 5)
		red.Toggle()
		time.Sleep(time.Second / 5)

	}

}
