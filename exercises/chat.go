package exercises

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"training.pl/examples/exercises/db"
)

const serverAddress = "localhost:8000"
const bufferSize = 128

func Chat() {
	if len(os.Args) > 1 {
		client()
	} else {
		server()
	}
}

type message struct {
	senderConn net.Conn
	bytes      []byte
}

func server() {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)
	log.Println("Listening on: " + serverAddress)

	connections := make([]net.Conn, 0)
	messages := make(chan *message, 100)
	mutex := &sync.RWMutex{}

	go func() {
		for msg := range messages {
			mutex.RLock()
			for _, conn := range connections {
				if conn != msg.senderConn {
					_, writeErr := conn.Write(msg.bytes)
					if writeErr != nil {
						log.Println("Write error: " + writeErr.Error())
					}
				}
			}
			mutex.RUnlock()
		}
	}()

	for {
		connection, acceptErr := listener.Accept()
		if acceptErr != nil {
			log.Println("Accept error: " + err.Error())
			continue
		}
		log.Println("Client connected: ", connection.LocalAddr())
		mutex.Lock()
		connections = append(connections, connection)
		mutex.Unlock()
		go handleConnection(connection, messages)
	}
	// close(messages)
}

func handleConnection(conn net.Conn, messages chan<- *message) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Close connection failed: " + err.Error())
		}
	}(conn)
	bytes := make([]byte, bufferSize)
	for {
		_, readErr := conn.Read(bytes)
		if readErr != nil {
			log.Println("Read error: " + readErr.Error())
			break
		}
		messages <- &message{conn, bytes}
	}
}

func client() {
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	defer func(connection net.Conn) {
		connectionErr := connection.Close()
		if connectionErr != nil {
			panic(connectionErr)
		}
	}(conn)

	go listenForMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		textBytes, _ := db.ToBytes(text)
		if len(textBytes) > bufferSize {
			log.Println("Message too long")
			continue
		}
		log.Printf("Sending message %d bytes long", len(textBytes))
		bytes := make([]byte, bufferSize)
		copy(bytes, textBytes[:])
		_, writeErr := conn.Write(bytes)
		if writeErr != nil {
			log.Println("Write error: " + writeErr.Error())
			break
		}
	}
}

func listenForMessages(conn net.Conn) {
	bytes := make([]byte, bufferSize)
	for {
		_, readErr := conn.Read(bytes)
		if readErr != nil {
			log.Println("Read error: " + readErr.Error())
			break
		}
		var text string
		err := db.FromBytes(bytes, &text)
		if err != nil {
			panic(err)
		}
		fmt.Println(text)
	}
}
