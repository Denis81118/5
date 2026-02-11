package main

import (
 "fmt"
 "time"
)

func main() {

 high := make(chan string, 5)
 low := make(chan string, 5)
 go func() {
  for {
   select {
   case task := <-high:
    fmt.Println(task)
   default:
    select {
    case task := <-low:
     fmt.Println( task)
    default:
     time.Sleep(1 * time.Second)
    }
   }
  }
 }()
 low <- "Фоновая задача 1"
 high <- "СРОЧНО!"
 low <- "Фоновая задача 2"
 high <- "ОЧЕНЬ ВАЖНО!"

 time.Sleep(2 * time.Second)
}