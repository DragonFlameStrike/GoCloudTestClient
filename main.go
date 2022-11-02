package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	CREATE = iota
	READ
	UPDATE
	DELETE
	EXIT
)

func main() {
	client := &http.Client{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n\nHello, it's client application for sends requests on server.\n" +
			"You should choose one of this option to continue\n" +
			"1. Create config on server (You should to have config on your PC to upload it on server)\n" +
			"2. Read config(-s)\n" +
			"3. Update config\n" +
			"4. Delete config\n" +
			"5. EXIT\n" +
			"Input number : ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		option, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Bad input? Try again")
			continue
		}
		option--

		switch option {
		case CREATE:
		case READ:
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
		case UPDATE:
		case DELETE:
		case EXIT:
			return
		default:
			fmt.Println("Bad input? Try again")
			continue
		}
		_, _ = reader.ReadString('\n')
	}
}
