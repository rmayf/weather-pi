package main 

import (
  "fmt"
  "github.com/rmayf/weather-pi/sensors"
)

func main() {
  t, e := dht11.Get(22)
  if e != nil {
    fmt.Println(e);
    return
  }
  fmt.Printf("The temperature is: %v\n", t)
}

