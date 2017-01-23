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
