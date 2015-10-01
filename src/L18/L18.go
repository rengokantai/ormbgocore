package main

import(
	"fmt"
	"net/http"
	"os"
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

func main(){
	urls:=[]string{"http://www.google.com/","http://facebook.com","http://yidi.me"}
	for _, url := range urls {
	pageLength, err := getPage(url)
	fmt.Printf("%d\n", pageLength)
	if err != nil{
		os.Exit(1)
	}
}
}