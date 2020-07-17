package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct {}

func main()  {
	res, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	//bs := make([]byte, 99999)
	//res.Body.Read(bs)
	//fmt.Println(string(bs))
	ls := LogWriter{}
	io.Copy(ls, res.Body)
}

func (LogWriter)Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The length of byte is ", len(bs))
	return len(bs), nil
}