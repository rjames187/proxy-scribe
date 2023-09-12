package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"proxy-scribe/handler"
	"strings"
)

var controller *handler.ProxyHandler

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("ProxyScribe > ")
	for reader.Scan() {
		input := parseInput(reader.Text())
		if len(input) < 1 {
			continue
		}
		switch input[0] {
		case "record":
			if controller != nil {
				fmt.Println("Proxy is already recording")
				break
			}
			go func () {
				log.Print("Proxy is listening")
				fmt.Println("ProxyScribe > ")
				controller = &handler.ProxyHandler{}
				log.Fatal(http.ListenAndServe(":4000", controller))
			}()
		case "finish":
			if controller == nil {
				fmt.Println("Cannot output spec because recording has not been started")
				break
			}
			controller.Doc.OutputSpec()
			fmt.Println("Spec has been outputted")
			os.Exit(0)
		default:
			fmt.Println("Command not found")
		}
		fmt.Print("ProxyScribe > ")
	}
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return strings.Split(input, " ")
}