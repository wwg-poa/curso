package main

import "fmt"

type Gopher struct {
  nome string
  cor  string
}

func main() {
  plush := Gopher{ nome: "Blu", cor: "azul" }
  fmt.Println(plush)
  fmt.Println(plush.nome)
  fmt.Println(plush.cor)
}
