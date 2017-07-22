package main

import (
  "fmt"
  "time"
)

func main() {
  go tarefa("goroutine")
  time.Sleep(time.Second)
  fmt.Println("meio")
  
  var entrada string
  fmt.Scanln(&entrada)
}

func tarefa(nome string) {
  fmt.Println("inicio", nome)
  time.Sleep(time.Second * 2)
  fmt.Println("fim", nome)
}
