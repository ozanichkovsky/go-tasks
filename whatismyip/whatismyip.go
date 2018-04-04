package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func main() {
	// TODO: read info from "http://httpbin.org/ip" and write origin IP to stdout
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	result := make(map[string]string)
	dec := json.NewDecoder(resp.Body)
	dec.Decode(&result)

	fmt.Println("My IP Address Is: ", result["origin"])
}
