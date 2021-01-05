This is a simple project I made to take some Adafruit Dotstar LED's I had lying around and turn them into entertaining lights on our Christmas tree.a

There is nothing particularly elegant about this code, this was just meant to be a quick and dirty project to make some fun patterns.a

I was running it on a Raspberry Pi Zero

```
GOARM=6 GOARCH=arm GOOS=linux go build -o rgbled main.go
```

I used a 74HCT245 level shifter/driver to convert from the Pi's 3.3V output to the LED's 5V