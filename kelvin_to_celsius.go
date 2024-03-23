package main

import "fmt"

const boilingPointKelvin float64 = 373.15 // Boiling point of water in Kelvin (100 degrees Celsius)

func main() {
  boilingPointCelsius := boilingPointKelvin - 273.15
  fmt.Printf("Boiling point of water in Celsius: %.2f°C\n", boilingPointCelsius)
}
