package main

import "fmt"

func main() {
  max_value := 100

  for i := 1; i <= max_value; i++ {
    if i % 3 == 0 && i % 5 == 0 {
      fmt.Println("pinpan"); continue
    }

    if i % 3 == 0 {
      fmt.Println("pin")
    }

    if i % 5 == 0 {
      fmt.Println("pan")
    }
  }
}
