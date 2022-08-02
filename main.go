package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	fmt.Println("Http multithreading testing tool.")
	fmt.Println("Usage: ./httptestgo.exe url concurrentCount repeattimes waitSecs_between_batches")
	fmt.Println("by Harry Xiao")
	fmt.Println("shawhu@gmail.com")
	fmt.Println("---------------------")

	argsWithoutProg := os.Args[1:]
	url := argsWithoutProg[0]
	if url == "" {
		//without url argument it won't run obviously
		return
	}
	//setting default values for variables
	concurrent, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		concurrent = 5
	}
	repeat, err := strconv.Atoi(argsWithoutProg[2])
	if err != nil {
		repeat = 1
	}
	batchwait, err := strconv.ParseFloat(argsWithoutProg[3], 64)
	if err != nil {
		batchwait = 0
	}
	//testing to see if the url is good to run
	fmt.Println("trying to access url: " + url)
	//try to access url
	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
		fmt.Println("Connection Error: " + err.Error())
		fmt.Println("exit with error")
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading http response Error: " + err.Error())
		fmt.Println("exit with error")
		return
	}
	_ = body

	fmt.Println("http get command returned")
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	fmt.Println("content:")
	//fmt.Println(string(body))

	fmt.Println("Now we tested the url, we can stress test it")
	fmt.Println("Starting the stress test")

	start := time.Now()
	ch := make(chan string)
	for j := 0; j < repeat; j++ {
		fmt.Println("Executing batch " + strconv.Itoa(j+1) + " x" + strconv.Itoa(concurrent))
		for i := 0; i < int(concurrent); i++ {
			go MakeRequest(url, ch)
		}
		time.Sleep(time.Duration(batchwait) * time.Second)
	}

	fmt.Printf("All done. Total requests count: %d. The stress test last %.2f seconds\n",
		int(concurrent)*repeat, time.Since(start).Seconds()-float64(repeat)*batchwait)
	fmt.Println("normal exit")
}
func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		fmt.Println("Connection Error: " + err.Error())
		fmt.Println("exit with error")
		return
	}
	defer resp.Body.Close()
	secs := time.Since(start).Seconds()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading http response Error: " + err.Error())
		fmt.Println("exit with error")
		return
	}
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}
