package main

import (
    "github.com/Markcial/pigeon"
    "fmt"
)

func doStuff(e *Pigeon.Event) {
    fmt.Print("yay!")
}

func main() {
  e := new(Pigeon.Event)
  o := new(Pigeon.Observer)
  o.Observe(e, doStuff)

  e.Fire()
  fmt.Printf("open your browser to this Url: http://localhost:8080")
  Pigeon.Main()
}