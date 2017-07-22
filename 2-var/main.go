package main

import "fmt"

func main() {
  var x int
  x = 10
  var y int = 20
  z := 30
  fmt.Println(x, y, z)
  
  nome := "Daniela"
  var n string = "Cacau"

  fmt.Println(nome)
  fmt.Println(n)

  verdade := true
  var mentira bool = false

  fmt.Println("Verdade é igual mentira:", verdade == mentira)

  java := 1.00 // float64
  var golang float64 = 2.50

  fmt.Println("java > go? ", java > golang)
}
