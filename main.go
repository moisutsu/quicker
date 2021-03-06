package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
		id  = flag.String("i", "", "id")
	)

	flag.Parse()

	if *msg == "" && *id == "" {
		fmt.Println("[Error] No argument.")
		os.Exit(1)
	}

	if *msg != "" && *id != "" {
		fmt.Println("[Error] To many arguments.")
		os.Exit(1)
	}

	baseURL := "https://quicker.ml/"

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
	} else {
		m := new(messageJSON)
		m.Message = *msg
		message, _ := json.Marshal(m)
		resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(message))
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
		var id idJSON
		if err := json.Unmarshal(b, &id); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(id.ID)
	}
}
