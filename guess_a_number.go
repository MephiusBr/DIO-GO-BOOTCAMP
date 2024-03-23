package main

import (
  "fmt"
  "math/rand"
  "strconv"
)

func main() {
  var secret_number int = get_random_number()
  fmt.Println("Let's start this round bro!.")

  start_round(&secret_number)
}

func start_round(chosen_number *int) {
  var guess_input string

  for {
    fmt.Printf("Guess the number >> ")
    fmt.Scanln(&guess_input)

    v, err := strconv.Atoi(guess_input)

    if err != nil {
      x := fmt.Errorf("%s", "you must enter an integer number between 1 and 100.")
      fmt.Println(x.Error())
      continue
    }

    if v == *chosen_number {
      fmt.Printf("You've entered %d, your input is correc!.\n", v)
      break
    }

    fmt.Println("Sorry, you got the wrong answer.")
  }

}

func get_random_number() int {
  var min, max int = 1, 100
  return rand.Intn(max - min) + min
}
