package main

import (
"fmt"
"time"
)

func emit(chanChannel chan chan string, done chan bool){
wordChannel := make(chan string)
chanChannel <- wordChannel

defer close(wordChannel)
words := []string{"The","quick","brown","fox"}
i:=0
t:= time.NewTimer(1 * time.Second)
for {
select{
case wordChannel <-words[i]:
i+=1
if i==len(words){
i=0
}
case<-done:
	done <-true
	return

case <- t.C:
	return
}

}

}


func main(){


channelch := make(chan chan string)
donech:= make(chan bool)

go emit(channelch, donech)
wch := <-channelch


for word:= range wch {
fmt.Print("%s ", word)
}




}