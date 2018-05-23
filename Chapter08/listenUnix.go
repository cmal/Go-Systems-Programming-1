package main

import (
	"fmt"
	"io"
	"net"
	_ "strconv"
	_ "time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, _ := r.Read(buf[:])
		fmt.Print("Read: ", string(buf[0:n]))
	}
}

func main() {

	l, err := net.Listen("unix", "./go.sock")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println("Accept error: ", err)
	}
	for {
		c.Write([]byte("hello..."))
		buf := make([]byte, 1024)
		for {
			n, _ := c.Read(buf[:])
			fmt.Print("Server Read: ", string(buf[0:n]))
		}

	}


}

