
package main

import(
"fmt"
"os"
)
func main(){
f,err := os.Open("test1.txt")
if err != nil{
fmt.Printf("%s\n", err)
os.Exit(1)
}

defer f.Close()

b:= make([]byte,100)

n,err:=f.Read(b)

str :=string(b)

fmt.Printf("%d; %c",n,str)
}
