package main

import (
    "bufio"
    "fmt"
    "os"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
		 
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
	
		    
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		fmt.Println(reader.ReadString('\n'))
	}()
	go func() {
		for i := 0; i < 1676; i++ {
			fmt.Println(<- c)
		}
		quit <- 0
	}()
	go func() {
		time.Sleep(8 * time.Second)
		quit <- 0
	}()
	fibonacci(c, quit)
}
