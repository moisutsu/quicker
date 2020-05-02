package main

import (
	"flag"
	"fmt"
	"os"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type messageJSON struct {
	Message string `json:"message"`
}

type idJSON struct {
	ID string `json:"id"`
}

func main() {
	var (
		msg = flag.String("m", "", "message")
		id = flag.String("i", "", "id")
	)

	flag.Parse()

	if *msg == "" && *id == "" {
		fmt.Println("[Error] No arguments.")
		os.Exit(1)
	}

	baseURL := "http://34.83.242.238/"

	if *msg == "" {
		values := url.Values{}
		values.Add("id", *id)
		resp, err := http.Get(baseURL + "?" + values.Encode())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		var message messageJSON
		if err := json.Unmarshal(b, &message); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(message.Message)
	}
}
