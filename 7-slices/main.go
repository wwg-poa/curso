package main

import "fmt"

func main() {

  linguagens := []string{ "python", "ruby", "java" }

  linguagens = append(linguagens, "go")

  fmt.Println("Linguagens:", linguagens)
  fmt.Println("Linguagens:", linguagens[1:3])

}
