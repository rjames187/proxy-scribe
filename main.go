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
			log.Print("Proxy is listening")
			controller := &handler.ProxyHandler{}
			log.Fatal(http.ListenAndServe(":4000", controller))
		}
		fmt.Print("ProxyScribe > ")
	}
}

func parseInput(input string) []string {
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return strings.Split(input, " ")
}