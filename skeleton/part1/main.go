// Solution to part 1 of the Whispering Gophers code lab.
// This program reads from standard input and writes JSON-encoded messages to
// standard output. For example, this input line:
//	Hello!
// Produces this output:
//	{"Body":"Hello!"}
//
package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Message struct {
	Body string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	enc := json.NewEncoder(os.Stdout)
	// TODO: Create a new bufio.Scanner reading from the standard input.
	// TODO: Create a new json.Encoder writing into the standard output.
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
