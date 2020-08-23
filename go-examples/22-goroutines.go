package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 5; i++ {
        fmt.Println(from, ":", i)
    }
}

func gor(from string, timer time.Duration){
  for i := 0; i < 5; i++ {
    fmt.Println(from, ":" , i)
    time.Sleep(timer)
  }
}

func main() {
    f("direct")

    go f("goroutine")

    go gor("goroutine2", 250)

    go gor("goroutine3", 750)

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}
