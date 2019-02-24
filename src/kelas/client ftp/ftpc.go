package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

const (
	uiDir  = "dir"
	uiCd   = "cd"
	uipwd  = "pwd"
	uiQuit = "quit"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host")
		os.Exit(1)
	}

	host := os.Args[1]

	conn, err := net.Dial("tcp", host+":1202")
	checkError(err)

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')

		line = strings.TrimRight(line, "\t\r\n")
		if err != nil {
			break
		}
		strs := strings.SplitN(line, " ", 2)

		switch strs[0] {
		case uiDir:
			dirRequest(conn)
		case uiCd:
			if len(strs) != 2 {
				fmt.Println("cd <dir>")
				continue
			}
			fmt.Println("CD \"", strs[1], "\"")
			cdRequest(conn, strs[1])
		case uipwd:
			pwdRequest(conn)
		case uiQuit:
			conn.Close()
			os.Exit(0)
		default:
			fmt.Println("Unknown Command!")
		}
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + ""))

	var buf [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, _ := conn.Read(buf[0:])
		result.Write(buf[0:n])
		length := result.Len()
		contents := result.Bytes()
		if string(contents[length-4:]) == "\r\n\r\n" {
			fmt.Println(string(contents[0 : length-4]))
			return
		}
	}
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " " + dir))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	if s != "OK" {
		fmt.Println("failed change dir")
	}
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD))
	var response [512]byte
	n, _ := conn.Read(response[0:])
	s := string(response[0:n])
	fmt.Println("current dir \"" + s + "\"")

}

func checkError(err error) {
	if err != nil {
		fmt.Println("fatal error", err.Error())
		os.Exit(1)
	}
}