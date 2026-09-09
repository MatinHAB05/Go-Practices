package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatalln("FATAL", err)
	}
}

var I = 0
var maxI = 20

func main() {
	IM, er := strconv.Atoi(os.Args[1])
	FatalIfError(er)
	conn, err := net.Dial("tcp", "127.97.97.97:8097")
	FatalIfError(err)

	defer func() {
		log.Println(conn.Close())
	}()
	log.Println("[[[[[[[[[[[[[[[[[	CLIENT	]]]]]]]]]]]]]]]]]")
	log.Println("conn.LocalAddr()", conn.LocalAddr())
	log.Println("conn.RemoteAddr()", conn.RemoteAddr(), "\n\n")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		log.Println("Client Read Routine")
		reader := bufio.NewReader(conn)
		for {
			mess, err := reader.ReadString(',')
			FatalIfError(err)
			log.Println("READ-MESSAGE:", mess)

		}

	}()

	go func() {
		log.Println("Client Write Routine")
		writer := bufio.NewWriter(conn)
		for {
			writer.WriteString(strconv.Itoa(IM) + "|" + strconv.Itoa(I) + ",")
			FatalIfError(err)
			writer.Flush()
			log.Println("WRITE-MESSAGE:", strconv.Itoa(I))
			time.Sleep(time.Second)

			I++
		}
	}()

	wg.Wait()

}
