package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("WEB REQUEST")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Printf("Respone type is : %T \n", response)

	data, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(data)
	fmt.Println(content)

}
