package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"./communicate"
)

//TODO: how to separate API logic, business logic and response format logic
func main() {
	fmt.Print("\nChoose Type of return data:\n1: JSON\n2: XML\n\nInput: ")
	scanner := bufio.NewScanner(os.Stdin)
	var isJSON bool
	for scanner.Scan() {
		text := scanner.Text()
		if text == "1" {
			isJSON = true
			break
		}
		if text == "2" {
			isJSON = false
			break
		}

		fmt.Print("\nPlease input 1 or 2")
		fmt.Print("\nInput: ")
	}

	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {
		communicate.HandleFunc(writer, request, isJSON)
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
