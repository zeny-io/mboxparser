package mboxparser

import (
	"github.com/blabber/mbox"
	"io"
	"os"
)

func Read(r io.Reader) (*Mbox, error) {
	var messages []*Message
	
	msgs := mbox.NewScanner(r)
	i := 0
	for msgs.Next() {
		messages = append(messages, Decode(msg.Message()))
		i++
	}

	return &Mbox{
		Messages: messages,
	}, nil
}

func ReadFile(filename string) (*Mbox, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return Read(fp)
}
