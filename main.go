package main

import (
	"bufio"
	"fmt"
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
			"3. Update config on server (You should to have config on your PC to update it on server)\n" +
			"4. Delete config\n" +
			"5. EXIT\n" +
			"Input number : ")
		text := readFromConsole(reader)
		option, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Bad input? Try again")
			continue
		}
		option--

		switch option {
		case CREATE:
			create(client, reader)
		case READ:
			read(client, reader)
		case UPDATE:
			edit(client, reader)
		case DELETE:
			delete(client, reader)
		case EXIT:
			return
		default:
			fmt.Println("Bad input? Try again")
			continue
		}
		_, _ = reader.ReadString('\n')
	}
}

func readFromConsole(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]
	if text[len(text)-1] == '\r' {
		text = text[:len(text)-1]
	}
	return text
}
