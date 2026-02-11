package main

import (	
	"fmt"
	"time"
)

func main(){
	
	done:=make(chan bool)
	go func(){
		for{
	select{
	case <- time.After(1*time.Second):
		fmt.Println("работает")
	case <- done:
		fmt.Println("все")
		return
		}
			}
				}()
	time.Sleep(3*time.Second)
	done <- true 
	time.Sleep(100*time.Millisecond)
	fmt.Println("завершена")
}

