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
	url:="http://www.google.com/"

	pageLength, err := getPage(url)
	fmt.Printf("%d", pageLength)
	if err != nil{
		os.Exit(1)
	}
}