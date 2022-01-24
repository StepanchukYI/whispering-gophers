// Solution to part 2 of the Whispering Gophers code lab.
//
// This program extends part 1.
//
// It makes a connection the host and port specified by the -dial flag, reads
// lines from standard input and writes JSON-encoded messages to the network
// connection.
//
// You can test this program by installing and running the dump program:
// 	$ go get github.com/campoy/whispering-gophers/util/dump
// 	$ dump -listen=localhost:8000
// And in another terminal session, run this program:
// 	$ part2 -dial=localhost:8000
// Lines typed in the second terminal should appear as JSON objects in the
// first terminal.
//
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"net"
	"os"
)

var dialAddr = flag.String("dial", "localhost:8000", "host:8000")

type Message struct {
	Body string
}

func main() {
	flag.Parse()
	// TODO: Parse the flags.

	conn, err := net.Dial("tcp", *dialAddr)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Open a new connection using the value of the "dial" flag.
	// TODO: Don't forget to check the error.

	scanner := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(conn)
	// TODO: Create a new bufio.Scanner reading from the standard input.
	// TODO: Create a json.Encoder writing into the connection you created before.
	for scanner.Scan() {
		x := scanner.Text()
		message := Message{Body: x}
		err := enc.Encode(message)
		if err != nil {
			log.Fatal(err)
		}
		// TODO: Create a new message with the read text.
		// TODO: Encode the message, and check for errors!
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// TODO: Check for a scan error.
}
