package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(err error, message string) {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main() {
	ln, err := net.Listen("tcp", ":80")
	check(err, "server is ready.")
	for {
		conn, err := ln.Accept()
		check(err, "Accepted Connection.")

		go func() {
			buf := bufio.NewReader(conn)

			for {
				name, err := buf.ReadString('\n')

				if err != nil {
					fmt.Printf("Client Disconect")
					break
				}
				conn.Write([]byte("Hello," + name))
			}
		}()
	}
}
