package main

import (
	"flag"
	"log"
	"net"
	"strconv"
)

func main() {
	port := flag.Int("port", 7777, "Port to accept connections on.")
	message := flag.String("message", "", "Message to send regardless of input.")
	flag.Parse()

	l, err := net.Listen("tcp", ":"+strconv.Itoa(*port))
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Listening to connections on port", strconv.Itoa(*port))
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Panicln(err)
		}

		go handleRequest(conn, *message)
	}
}

func handleRequest(conn net.Conn, message string) {
	log.Println("Accepted new connection.")
	defer conn.Close()
	defer log.Println("Closed connection.")

	if len(message) > 0 {
		conn.Write([]byte(message + "\n"))
		return
	}

	buf := make([]byte, 1024)
	size, err := conn.Read(buf)
	if err != nil {
		return
	}
	data := buf[:size]
	log.Println("Read new data from connection", data)
	conn.Write(data)
}
