package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatalln("FATAL", err)
	}
}

func MAINFUNCTION(server net.Listener) {
	conn, err := server.Accept()
	FatalIfError(err)
	go MAINFUNCTION(server)

	log.Println("[[[[[[[[[[[[[[[[[	SERVER	]]]]]]]]]]]]]]]]]")
	log.Println("conn.LocalAddr()", conn.LocalAddr())
	log.Println("conn.RemoteAddr()", conn.RemoteAddr(), "\n\n")

	go func() {
		log.Println("Server Read Routine")
		reader := bufio.NewReader(conn)
		for {
			mess, err := reader.ReadString(',')
			FatalIfError(err)
			log.Println("READ-MESSAGE:", mess)

		}

	}()

	go func() {
		log.Println("Server Write Routine")
		writer := bufio.NewWriter(conn)
		for {
			writer.WriteString("HelloWorldThisIsMatinHasanaliBaki138397.com" + ",")
			FatalIfError(err)
			writer.Flush()
			log.Println("WRITE-MESSAGE:", "HelloWorldThisIsMatinHasanaliBaki138397.com")
			time.Sleep(time.Second)
		}
	}()
	MAINFUNCTION(server)
}

func main() {
	server, err := net.Listen("tcp", "127.97.97.97:8097")
	FatalIfError(err)
	log.Println("LISTNEING...")
	MAINFUNCTION(server)
}
