package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

// Сетевой адрес.
const addr = "0.0.0.0:12345"

// Протокол сетевой службы.
const protocol = "tcp4"

// Поговорки GO.
var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

func randomProverb() string {
	rand.Seed(time.Now().UnixNano())
	return proverbs[rand.Intn(len(proverbs))]
}

func handleConn(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	for {
		proverb := randomProverb() + "\r\n"
		_, err := conn.Write([]byte(proverb))
		if err != nil {
			log.Println("Error writing to client:", err)
			return
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {

	listener, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConn(conn)
		log.Println("Server started, waiting for connections...")
	}
}
