package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(k8sv1.NamespaceDefault)
	// rest.InClusterConfig()
	// io.Copy(dst, src)
	go func() {
		fmt.Println("hello world-1")
		time.Sleep(time.Second)
		return
	}()
	fmt.Println("hello world - 2")
}
