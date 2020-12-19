package pattern

import (
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
	purple = rpioapa102.LED{Red: 255, Green: 0, Blue: 255, Brightness: 31}
	cyan   = rpioapa102.LED{Red: 0, Green: 255, Blue: 255, Brightness: 31}
	yellow = rpioapa102.LED{Red: 255, Green: 255, Blue: 0, Brightness: 31}
)

type Chase struct{}

func (Chase) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		l[i] = off
	}
	c.Write(l)

	colors := []rpioapa102.LED{red, green, blue, purple, cyan, yellow}

	for cl := 0; cl < len(colors); cl++ {
		for i := 0; i < len(l); i++ {
			if i > 0 {
				l[i-1] = off
			}
			l[i] = colors[cl]
			c.Write(l)
			time.Sleep(10 * time.Millisecond)
		}

		for i := len(l) - 1; i >= 0; i-- {
			if i < len(l)-1 {
				l[i+1] = off
			}
			l[i] = colors[cl]
			c.Write(l)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

type Collide struct{}

func (Collide) Display(c rpioapa102.LEDController, l []rpioapa102.LED) {
	for i := 0; i < len(l); i++ {
		if i > 0 {
			l[i-1] = off
			l[len(l)-i] = off
		}
		if i > len(l)/2 {
			l[i] = purple
			l[len(l)-1-i] = purple
		} else {
			l[i] = red
			l[len(l)-1-i] = blue
		}
		c.Write(l)
		time.Sleep(10 * time.Millisecond)
	}
}
