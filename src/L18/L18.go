package main

import(
	"fmt"
	"net/http"
	//"os"
	"io/ioutil"
)

func getPage(url string)(int,error){
	resp, err := http.Get(url)
	if err!=nil{
		return 0, err
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return 0 ,err
	}

	return len(body),nil
}


func getter(url string, size chan int){
	length,err :=getPage(url)
	if err ==nil{
		size<-length
	}
	
}

func main(){
	urls:=[]string{"http://www.google.com/","http://facebook.com","http://yidi.me"}

	size:=make(chan int)
	for _, url := range urls {
		go getter(url, size)
	// pageLength, err := getPage(url)
	// fmt.Printf("%d\n", pageLength)
	// if err != nil{
	// 	os.Exit(1)
	// }
		}
	for i:=0; i<len(urls);i++{
		fmt.Printf("%d",<-size)
	}

}