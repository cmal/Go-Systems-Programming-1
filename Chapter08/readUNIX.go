package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, _ := r.Read(buf[:])
		fmt.Print("Read: ", string(buf[0:n]))
	}
}

func main() {

	c, err := net.Dial("unix", "./go.sock")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	go readSocket(c)
	n := 0
	for {
		message := []byte("Hi there: " + strconv.Itoa(n) + "\n")
		_, _ = c.Write(message)
		time.Sleep(5 * time.Second)
		n = n + 1
	}
}
