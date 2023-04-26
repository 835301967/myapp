package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main()  {
	start := time.Now()
	inputChan :=make(chan string)
	for _,url:=range os.Args[1:]{
		if !strings.HasPrefix(url,"http://"){
			url = "http://"+url
		}
		go fetch(url,inputChan)
	}
	for range os.Args[1:]{
		fmt.Printf("%s\n",<-inputChan)
	}
	fmt.Printf("%.2fs\n",time.Since(start).Seconds())
}

func fetch(url string,inputChan chan<- string)  {
	start := time.Now()
	response, err := http.Get(url)
	if err!=nil{
		inputChan <- fmt.Sprintln(err)
		return
	}
	written, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close()
	if err!=nil{
		inputChan<- fmt.Sprintln(err)
		return
	}
	sec:=time.Since(start).Seconds()
	inputChan <- fmt.Sprintf("%.2f  %7d  %s ----\n",sec,written,url)

}
