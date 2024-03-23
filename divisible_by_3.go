package main

import "fmt"

func main() {
  var initial_value, max_value int = 1, 100

  for i := initial_value; i <= max_value; i++ {
    if i % 3 == 0 {
      fmt.Printf("Number %d is divisible by 3.\n", i)
    }
  }
}
