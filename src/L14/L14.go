package main

import(
"fmt"
"os"
)
func printer(msg string) error {
_, err:= fmt.Print("%s\n",msg)
return err
}
func main(){
if err:=printer("ERROR"); err !=nil{
os.Exit(1)
}}