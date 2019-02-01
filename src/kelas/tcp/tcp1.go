package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide port number")
		os.Exit(100)
	}
	port := ":" + arguments[1]
	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		fmt.Print("-> ", string(netData))
		c.Write([]byte(netData))
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("exiting tcp server!")
			return
		}
	}
}
