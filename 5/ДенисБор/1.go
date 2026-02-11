package main

import (	
	"fmt"
	"time"
)

func main(){
	ch := make(chan string)

	go func(){
		fmt.Println("Начинаем запрос")
		time.Sleep(3*time.Second)
		ch <- "всё"
	}()
	select {
	case result := <-ch:
		fmt.Println("Запрос завершен до таймера", result)
		case <- time.After(2*time.Second):
		fmt.Println("Таймаут 2 секунды")
	}
	fmt.Println("Закончили")
}