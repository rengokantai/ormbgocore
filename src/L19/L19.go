package main

import(
	"fmt"
	"time"
)

func printer(msg string, stopCh chan bool){
	for{
		select {
			case <-stopCh:
			return
		default:
				fmt.Printf("%s",msg)
		}
	}

	}



func main(){
	//goCh := make(chan bool)
	stopCh := make(chan bool)
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer %d",i), stopCh)
	}
	time.Sleep(5* time.Second)
	close(stopCh)
	time.Sleep(5*time.Second)
}