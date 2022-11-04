package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func create(client *http.Client, reader *bufio.Reader) {
	fmt.Println("\n You chose create option.\n" +
		"You should to input filename like - data.json\n" +
		"Or input filepath with filename like - ./content/data.json\n" +
		"file shouldn't contain version\n" +
		"Filename: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", text)
	if err != nil {
		return
	}
	file, err := os.Open(text)
	if err != nil {
		return
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return
	}
	writer.Close()
	req, err := http.NewRequest("POST", "http://localhost:8080/config", bytes.NewReader(body.Bytes()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return
}

func read(client *http.Client, reader *bufio.Reader) {
	fmt.Println("\n You chose read option.\n" +
		"If you want to see all configs on the server - press enter\n" +
		"If you want to see config by service name - you should to input query like \"service=kuber\"")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	query := ""
	if text != "" {
		query += "?" + text
	}
	req, _ := http.NewRequest("GET", "http://localhost:8080/config"+query, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}