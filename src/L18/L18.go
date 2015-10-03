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


// func getter(url string, size chan string){
// 	length,err :=getPage(url)
// 	if err ==nil{
// 		size<-fmt.Sprintf("%d\n",length)
// 	}
	
// }
func worker(urlCh chan string, sizeCh chan string){
	for {
		url:=<-urlCh
		length,err:= getPage(url)
	
	if err==nil {
		sizeCh <- fmt.Sprintf("%d",length)
	}else{
		sizeCh<-fmt.Sprintf("Error")
	}
}
}

func generator(url string,urlCh chan string){
	urlCh<-url
}
func main(){
	

	urlCh :=make(chan string)
	sizeCh :=make(chan string)



	for i := 0; i < 10; i++ {
			go worker(urlCh,sizeCh)
	}

	urls:=[]string{"http://www.google.com/","http://facebook.com","http://yidi.me"}


	for _, url :=range urls {
		go generator(url,urlCh)
		// or urlCh<-url
	}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n",<-sizeCh)
	}







/*	size:=make(chan string)
	for _, url := range urls {
		go getter(url, size)
	// pageLength, err := getPage(url)
	// fmt.Printf("%d\n", pageLength)
	// if err != nil{
	// 	os.Exit(1)
	// }
		}
	for i:=0; i<len(urls);i++{
		fmt.Printf("%s",<-size)
	}
*/
}