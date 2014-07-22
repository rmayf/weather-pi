package dht11 

import (
  "github.com/davecheney/gpio"
  "github.com/davecheney/gpio/rpi"
  "fmt"
  "time"
)

// Get returns the temperature data of the DHT11 IC
func Get(pin_num int) (t int, e error) {
  switch {
  case pin_num == 21:
    pin_num = rpi.GPIO21
  case pin_num == 22:
    pin_num = rpi.GPIO22
  case pin_num == 23:
    pin_num = rpi.GPIO23
  case pin_num == 24:
    pin_num = rpi.GPIO24
  case pin_num == 27:
    pin_num = rpi.GPIO27
  case pin_num == 17:
    pin_num = rpi.GPIO17
  default:
    return 0, fmt.Errorf("temp: invalid GPIO number %d", pin_num)
  }

  pin, e := gpio.OpenPin(pin_num, gpio.ModeOutput)
  if e != nil {
    return
  }
  // Reset the line
  pin.Set();
  time.Sleep(500 * time.Millisecond)
  pin.Clear();
  time.Sleep(20 * time.Millisecond)
  pin.Set();
  // Wait for start signal from the sensor
  pin.SetMode(gpio.ModeInput)
  t = 0
  pin.BeginWatch(gpio.EdgeFalling, func() {
    t++
  })
  time.Sleep(5 * time.Second)
  pin.Close();
  return
}
