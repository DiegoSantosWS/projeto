package main

import (
	"fmt"
	"log"

	"github.com/ua-parser/uap-go/uaparser"
)

func main() {
	uagent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.167 Safari/537.36"

	parser, err := uaparser.New("regexes.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	client := parser.Parse(uagent)

	fmt.Println("BROWSER: ", client.UserAgent.Family) // "Amazon Silk"
	fmt.Println("BYTE: ", client.UserAgent.Major)     // "1"
	fmt.Println("NÃO SEI: ", client.UserAgent.Minor)  // "1"
	fmt.Println("NÃO SEI: ", client.UserAgent.Patch)  // "0-80"
	fmt.Println("OS: ", client.Os.Family)             // "Android"
	fmt.Println(client.Os.Major)                      // ""
	fmt.Println(client.Os.Minor)                      // ""
	fmt.Println(client.Os.Patch)                      // ""
	fmt.Println(client.Os.PatchMinor)                 // ""
	fmt.Println(client.Device.Family)                 // "Kindle Fire"

}
