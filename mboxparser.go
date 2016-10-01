package mboxparser

import (
	"github.com/bthomson/mbox"
	"io"
	"os"
)

func Read(r io.Reader) (*Mbox, error) {
	msgs, err := mbox.Read(r, false)
	if err != nil {
		return nil, err
	}

	messages := make([]*Message, len(msgs))
	for i, msg := range msgs {
		messages[i] = Decode(msg)
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
