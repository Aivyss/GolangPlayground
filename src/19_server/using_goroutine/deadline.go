package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
		}

		go handleWithDeadline(conn)
	}
}

func handleWithDeadline(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second)) // the deadline of tcp connection is now + 10 second.
	if err != nil {
		log.Println("CONNECTION TIMEOUT!")
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println("request message: ", ln)
		fmt.Fprint(conn, "request message: ", ln+"\n")
	}

	fmt.Println("you get here by connection deadline.")
}
