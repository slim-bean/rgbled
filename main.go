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

	d := pattern.Chase{}
	c := pattern.Collide{}
	b := pattern.Breath{}
	for {
		if time.Now().Hour() > 5 {
			d.Display(controller, rgbLEDs)
			c.Display(controller, rgbLEDs)
			c.Display(controller, rgbLEDs)
			b.Display(controller, rgbLEDs)
		} else {
			for i := 0; i < len(rgbLEDs); i++ {
				rgbLEDs[i] = rpioapa102.LED{0, 0, 0, 0}
			}
			controller.Write(rgbLEDs)
			time.Sleep(10 * time.Second)
		}
	}
}

func rando(d []rpioapa102.LED) {
	for i := 0; i < len(d); i++ {
		d[i] = rpioapa102.LED{
			Red:        randomByte(),
			Green:      randomByte(),
			Blue:       randomByte(),
			Brightness: 10,
		}
	}
}

func randomByte() byte {
	return byte(random.Intn(256))
}
