package mboxparser

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestNotFound(t *testing.T) {
	_, err := ReadFile("testdata/not_found.mbox")
	if err == nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	mbox, err := ReadFile("testdata/simple.mbox")
	if err != nil {
		t.Fatal(err)
	}

	if len(mbox.Messages) != 1 {
		t.Fatalf("invalid messages count: %d", len(mbox.Messages))
	}
}

func TestNoBody(t *testing.T) {
	mbox, err := ReadFile("testdata/nobody.mbox")
	if err != nil {
		t.Fatal(err)
	}

	message := mbox.Messages[0]
	if len(message.Bodies) != 0 {
		t.Fatal("Invalid body found")
	}
}

func TestPlain(t *testing.T) {
	mbox, err := ReadFile("testdata/plain.mbox")
	if err != nil {
		t.Fatal(err)
	}

	message := mbox.Messages[0]
	if len(message.Bodies) != 1 {
		t.Fatal("Invalid body found")
	}

	body := message.Bodies[0]
	text, err := ioutil.ReadAll(body.Content)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(text)) != "This is a test" {
		t.Fatalf("Invalid body: %s", string(text))
	}
}
