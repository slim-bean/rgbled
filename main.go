package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	rpioapa102 "github.com/aggronerd/rpio-apa102"
	"github.com/stianeikeland/go-rpio/v4"

	"rgbled/pkg/pattern"
)

const ledCount = 252

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	starPin := rpio.Pin(17)
	starPin.Mode(rpio.Output)
	starPin.Low()
	starOn := false

	controller := rpioapa102.NewLEDController(rpio.Spi0)

	rgbLEDs := make([]rpioapa102.LED, ledCount)

	// catch signals and terminate the app
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	// monitor for signals in the background
	go func() {
		s := <-sigc
		//ticker.Stop()
		fmt.Println("\nreceived signal:", s)
		//time.Sleep(speed)
		// turn off the LEDs
		for i := 0; i < len(rgbLEDs); i++ {
			rgbLEDs[i] = rpioapa102.LED{1, 1, 1, 0}
		}
		controller.Write(rgbLEDs)
		controller.Finish()
		time.Sleep(400 * time.Millisecond)
		os.Exit(0)
	}()

	cr := pattern.RedChase{}
	cg := pattern.GreenChase{}
	cb := pattern.BlueChase{}
	cc := pattern.CyanChase{}
	cv := pattern.VioletChase{}
	c := pattern.Collide{}
	b := pattern.Breath{}
	nye := pattern.NewYear{}

	d := []pattern.Pattern{cr, cg, cb, cc, cv, c, c, c, b, b, b}

	for {
		if time.Now().Month() == time.December &&
			time.Now().Day() == 31 &&
			time.Now().Hour() == 23 &&
			time.Now().Minute() == 59 &&
			time.Now().Second() > 49 {
			nye.Display(controller, rgbLEDs)
			for i := 0; i < 100; i++ {
				d[random.Intn(len(d)-1)].Display(controller, rgbLEDs)
			}
		} else if time.Now().Hour() > 5 {
			if !starOn {
				starPin.High()
				starOn = true
			}
			d[random.Intn(len(d)-1)].Display(controller, rgbLEDs)
		} else {
			if starOn {
				starPin.Low()
				starOn = false
			}
			for i := 0; i < len(rgbLEDs); i++ {
				rgbLEDs[i] = rpioapa102.LED{0, 0, 0, 0}
			}
			controller.Write(rgbLEDs)
			time.Sleep(10 * time.Second)
		}
	}
}
