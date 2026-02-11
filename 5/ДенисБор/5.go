package main

import (
 "fmt"
 "time"
)

func main() {
 stop := make(chan bool)
 
 go func() {
  for {
   select {
   case <-stop:
    fmt.Println("завершается")
    return
   default:
    fmt.Println("работает")
    time.Sleep(400 * time.Millisecond)
   }
  }
 }()

 time.Sleep(2 * time.Second)
 stop <- true
 time.Sleep(100 * time.Millisecond)



  go func() {
  for {
   select {
   case <-stop:
    fmt.Println("2 горутина завершается")
    return
   default:
    fmt.Println(" 2 горутина работает")
    time.Sleep(400 * time.Millisecond)
   }
  }
 }()

 time.Sleep(2 * time.Second)
 stop <- true
 time.Sleep(100 * time.Millisecond)
}