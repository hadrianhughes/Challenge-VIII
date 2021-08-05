package main

import (
  "fmt"
  "os"
)

var charValues = map[byte]int{
  'M': 1000,
  'D': 500,
  'C': 100,
  'L': 50,
  'X': 10,
  'V': 5,
  'I': 1,
}

func numeralToValue(numeral string) (value int) {
  value = 0
  ones := 0

  for i := 0; i < len(numeral); i++ {
    if numeral[i] == 'I' {
      ones++
    } else {
      if ones > 1 {
        fmt.Println("Only 1 I allowed for subtractive notation")
        return
      }

      value += (charValues[numeral[i]] - ones)
      ones = 0
    }
  }

  value += ones

  return
}

func main() {
  args := os.Args

  if len(args) < 3 {
    fmt.Println("Requires 2 arguments")
    return
  }

  a := numeralToValue(args[1])
  b := numeralToValue(args[2])

  fmt.Println(a < b)
}