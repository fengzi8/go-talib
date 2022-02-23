package main

import (
    "bufio"
    "fmt"
    "os"
)
func main() {

    messages := make(chan string)

    go func() { 
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter text: ")
		text, _ := reader.ReadString('\n')
		text2 := ""
		fmt.Scanln(&text2)
		messages <- text2 
		messages <- text
		close(messages) }()

    msg := <-messages
    fmt.Println(msg)
	msg2 := <-messages
    fmt.Println(msg2)
    // reader := bufio.NewReader(os.Stdin)
    // fmt.Print("Enter text: ")
    // text, _ := reader.ReadString('\n')909

    // fmt.Println(text)

    // fmt.Println("Enter text: ")
    // text2 := ""
    // fmt.Scanln(&text2)
    // fmt.Println(text2)

    // ln := ""
    // fmt.Sscanln("%v", ln)
    // fmt.Println(ln)


	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text: ")
	// text, _ := reader.ReadString('\n')
	// fmt.Println(text)
	// select {}
}