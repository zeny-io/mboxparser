package mboxparser

import (
	"github.com/blabber/mbox"
	"io"
	"os"
)

func Read(r io.Reader) (*Mbox, error) {
	var messages []*Message
	
	msgs := mbox.NewScanner(r)
	buf := make([]byte, 0, 64*1024)
	msgs.Buffer(buf, 1024*1024*100)
	for msgs.Next() {
		messages = append(messages, Decode(msg.Message()))
	}
	// Check for an error, if so report that error.
	if msgs.Err() != nil {
		fmt.Printf("Possible error reading mbox:\n%v\n", msgs.Err())
	}

	return &Mbox{
		Messages: messages,
	}, msgs.Err()
}

func ReadFile(filename string) (*Mbox, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return Read(fp)
}
