package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	servicehandler "thien.com/service-handler"
)

//TODO: how to separate API logic, business logic and response format logic
func main() {

	fmt.Print("\tPlease select render type:\n1: JSON\n2: XML\n\nInput: ")
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

	http.HandleFunc("/postWithComments", servicehandler.Get(isJSON))

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
