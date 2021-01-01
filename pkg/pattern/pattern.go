package pattern

import (
	"math/rand"
	"time"

	rpioapa102 "github.com/aggronerd/rpio-apa102"
)

type Pattern interface {
	Display(c rpioapa102.LEDController, l []rpioapa102.LED)
}

var (
	off    = rpioapa102.LED{Red: 0, Green: 0, Blue: 0, Brightness: 0}
	red    = rpioapa102.LED{Red: 255, Green: 0, Blue: 0, Brightness: 31}
	green  = rpioapa102.LED{Red: 0, Green: 255, Blue: 0, Brightness: 31}
	blue   = rpioapa102.LED{Red: 0, Green: 0, Blue: 255, Brightness: 31}
	violet = rpioapa102.LED{Red: 255, Green: 0, Blue: 255, Brightness: 31}
	cyan   = rpioapa102.LED{Red: 0, Green: 255, Blue: 255, Brightness: 31}
	yellow = rpioapa102.LED{Red: 255, Green: 255, Blue: 0, Brightness: 31}
	white  = rpioapa102.LED{Red: 255, Green: 255, Blue: 255, Brightness: 31}
)

var (
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type RedChase struct{}

func (RedChase) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = off
	}
	c.Write(l)

	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
		}
		l[i] = red
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for i := len(l) - 1; i >= 0; i-- {
		if i < len(l)-6 {
			l[i+5] = off
		}
		l[i] = red
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type GreenChase struct{}

func (GreenChase) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = off
	}
	c.Write(l)

	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
		}
		l[i] = green
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for i := len(l) - 1; i >= 0; i-- {
		if i < len(l)-6 {
			l[i+5] = off
		}
		l[i] = green
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type BlueChase struct{}

func (BlueChase) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = off
	}
	c.Write(l)

	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
		}
		l[i] = blue
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for i := len(l) - 1; i >= 0; i-- {
		if i < len(l)-6 {
			l[i+5] = off
		}
		l[i] = blue
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type CyanChase struct{}

func (CyanChase) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = off
	}
	c.Write(l)

	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
		}
		l[i] = cyan
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for i := len(l) - 1; i >= 0; i-- {
		if i < len(l)-6 {
			l[i+5] = off
		}
		l[i] = cyan
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type VioletChase struct{}

func (VioletChase) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = off
	}
	c.Write(l)

	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
		}
		l[i] = violet
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for i := len(l) - 1; i >= 0; i-- {
		if i < len(l)-6 {
			l[i+5] = off
		}
		l[i] = violet
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type Collide struct{}

func (Collide) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
			l[len(l)-i+5] = off
		}
		if i > len(l)/2 {
			l[i] = violet
			l[len(l)-1-i] = violet
		} else {
			l[i] = red
			l[len(l)-1-i] = blue
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < len(l); i++ {
		if i > 5 {
			l[i-6] = off
			l[len(l)-i+5] = off
		}
		if i > len(l)/2 {
			l[i] = cyan
			l[len(l)-1-i] = cyan
		} else {
			l[i] = blue
			l[len(l)-1-i] = green
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type Breath struct{}

func (Breath) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < 256; i += 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: byte(i), Green: 0, Blue: 0, Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 255; i >= 0; i -= 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: byte(i), Green: 0, Blue: 0, Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < 256; i += 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: 0, Green: byte(i), Blue: 0, Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 255; i >= 0; i -= 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: 0, Green: byte(i), Blue: 0, Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < 256; i += 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: 0, Green: 0, Blue: byte(i), Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 255; i >= 0; i -= 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: 0, Green: 0, Blue: byte(i), Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < 256; i += 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: 0, Green: byte(i), Blue: byte(i), Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 255; i >= 0; i -= 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: 0, Green: byte(i), Blue: byte(i), Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 0; i < 256; i += 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: byte(i), Green: 0, Blue: byte(i), Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for i := 255; i >= 0; i -= 5 {
		for j := 0; j < len(l); j++ {
			l[j] = rpioapa102.LED{Red: byte(i), Green: 0, Blue: byte(i), Brightness: 31}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}

type NewYear struct{}

func (NewYear) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = red
	}
	c.Write(l)

	scale := float64(10*time.Second.Milliseconds()) / float64(len(l))
	color := red
	for i := 10 * time.Second.Milliseconds(); i >= 0; i -= 10 {
		if i == 9*time.Second.Milliseconds() {
			color = rpioapa102.LED{Red: 255, Green: 255, Blue: 128, Brightness: 31}
		} else if i == 8*time.Second.Milliseconds() {
			color = rpioapa102.LED{Red: 255, Green: 128, Blue: 255, Brightness: 31}
		} else if i == 7*time.Second.Milliseconds() {
			color = rpioapa102.LED{Red: 128, Green: 255, Blue: 255, Brightness: 31}
		} else if i == 6*time.Second.Milliseconds() {
			color = violet
		} else if i == 5*time.Second.Milliseconds() {
			color = cyan
		} else if i == 4*time.Second.Milliseconds() {
			color = blue
		} else if i == 3*time.Second.Milliseconds() {
			color = green
		} else if i == 2*time.Second.Milliseconds() {
			color = yellow
		} else if i == 1*time.Second.Milliseconds() {
			color = red
		}

		led := int(float64(10*time.Second.Milliseconds()-i) / scale)

		for j := 0; j < len(l); j++ {
			if j < led {
				if j > 5 {
					l[j-6] = off
				}
			} else {
				l[j] = color
			}
		}

		time.Sleep(10 * time.Millisecond)
		c.Write(l)
	}

	for k := 0; k < 100; k++ {
		for i := 0; i < len(l); i++ {
			l[i] = white
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
		for i := 0; i < len(l); i++ {
			l[i] = off
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for k := 0; k < 500; k++ {
		for i := 0; i < len(l); i++ {
			l[i] = rpioapa102.LED{
				Red:        randomByte(),
				Green:      randomByte(),
				Blue:       randomByte(),
				Brightness: 31,
			}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
	for k := 0; k < 100; k++ {
		for i := 0; i < len(l); i++ {
			l[i] = white
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
		for i := 0; i < len(l); i++ {
			l[i] = off
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

	for k := 0; k < 500; k++ {
		for i := 0; i < len(l); i++ {
			l[i] = rpioapa102.LED{
				Red:        randomByte(),
				Green:      randomByte(),
				Blue:       randomByte(),
				Brightness: 31,
			}
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}

}

func randomByte() byte {
	return byte(random.Intn(256))
}
