package main

import (
	"os"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {
	arg := os.Args[1]

	resp, err := http.Get(arg)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		fmt.Print(bodyString)
	}
}
