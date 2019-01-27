package main

import ("fmt"
		"net"
		"bytes"
		"io"
		"os"
	)

const myIPAddress = "192.168.60.138"
const ipv4HeaderSize = 20

func main()  {
	if len(os.Args) != 2 {
		fmt.Fprint(os.Stderr, "Usage:",os.Args[0],"host") 
		os.Exit(1)
	}
	localAddr, err := net.ResolveIPAddr("ip4",myIPAddress)

	if err != nil {
		fmt.Println("Resolution error" , err.Error())
		os.Exit(1)
	}
	RemoteAddr, err := net.ResolveIPAddr("ip4",os.Args[1])
	if err != nil {
		fmt.Println("Resolution error" , err.Error())
		os.Exit(1)
	}
	conn,err := net.DialIP("ip4:icmp",localAddr, RemoteAddr)
	checkError(err)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)


		_,err = conn.Write(msg[0:len])
		checkError(err)

		fmt.Println("message sent:	")
		for n := 0; n < 8; n++ {
			fmt.Println(" ", msg[n])
		}
		fmt.Println()


		size,err2 := conn.Read(msg[0:])
		checkError(err2)

		fmt.Println("message recived:	")
		for n := ipv4HeaderSize; n < size; n++ {
			fmt.Print(" ",msg[n])
		}
		fmt.Println()
		os.Exit(0)
}

func checkSum(msg []byte) uint16 {
	sum := 0

	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}


func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error() )
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for{
		n,err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil,err
		}
	}
	return result.Bytes(),nil
}
