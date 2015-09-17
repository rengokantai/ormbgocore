package main

import (
"fmt"
)

func emit(c chan string){

words := []string{"The","quick","brown","fox"}


for _,word := range words{


c<- word
}

close(c)

}


func main(){

wordChannel :=make(chan string)

go emit(wordChannel)

//for word := range wordChannel{
//fmt.Printf("%s", word)
//}

word:= <- wordChannel
fmt.Printf("%s", word)
word2, ok:= <- wordChannel
fmt.Printf("%s", word2, ok)
word3, ok:= <- wordChannel
fmt.Printf("%s", word3, ok)
word4, ok:= <- wordChannel
fmt.Printf("%s", word4, ok)
word5, ok:= <- wordChannel
fmt.Printf("%s", word5, ok)
}