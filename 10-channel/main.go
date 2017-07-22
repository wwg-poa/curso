package main

import (
  "fmt"
  "time"
)

func main() {
  canal := make(chan int)

  go tarefa("goroutine", canal)

  for i := 0; i < 100; i++ {
    canal <- i
  }

  time.Sleep(time.Second * 3)
}

func tarefa(nome string, canal chan int) {
  for {
   j := <-canal
   fmt.Println(nome, j)
   time.Sleep(time.Millisecond * 100)
  }
}
