package main

import (	
	"fmt"
	
)

func main(){
	ch:= make(chan string, 1)
	ch <- "привет"//если это убрать выведет "данных нет"
	select {
	case message := <- ch:
		fmt.Println("получено: ", message)
	default:
		fmt.Println("данных нет")
	}
}