package main
import (
"fmt"
"math/rand"
)

func emit(c chan int){

for{
c<-rand.Intn(100)
}

}


func main(){

randoms :=make(chan int)

go emit(randoms)

for word := range randoms{
fmt.Printf("%d", word)
}


}