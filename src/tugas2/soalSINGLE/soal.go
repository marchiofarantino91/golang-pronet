package main

import ("fmt"
		"net"
		"bufio"
)
func check(err error, message string)  {
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", message)
}

func main()  {
	ln, err := net.Listen("tcp", ":8080")
	check(err,"server is ready.")
	for{
		conn,err := ln.Accept()
		check(err, "accepted connection.")

		go func(){
			buf := bufio.NewReader(conn)

			for{
				name,err := buf.ReadString('\n')

				if err != nil{
					fmt.Printf("Client disconect")
					break
				}
				conn.Write([]byte("hello," +name))
			}
		}()
	}
}