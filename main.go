package main

import (
	"fmt"
	"io"
	"os"
)

type User struct {
	Name string
	Addr string
}

func main() {
	// waitInterrupt := make(chan os.Signal, 1)
	// signal.Notify(waitInterrupt, os.Interrupt, syscall.SIGTERM)
	//
	// runningChan := make(chan error)
	//
	// fmt.Println("starting......")
	// go func() {
	// 	runningChan <- nil
	// }()
	// select {
	// case <-waitInterrupt:
	// 	// Make a new line in the terminal
	// 	fmt.Println()
	// 	return
	// case err := <-runningChan:
	// 	if err != nil {
	// 		return
	// 		// fmt.Println("return")
	// 	}
	// }
	//
	// fmt.Println("hello world")

	r, w := io.Pipe()
	go func() {
		fmt.Fprint(w, "hello world")
		w.Close()
	}()
	io.Copy(os.Stdout, r)
}
