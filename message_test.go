package mboxparser

import (
	"testing"
)

func TestHeader(t *testing.T) {
	header := make(Header)

	testHeader := "qwertyuiop"
	header.Set("x-mBox-header", testHeader)

	retval := header.Get("x-mbox-HeadeR")
	if retval != testHeader {
		t.Fatalf("Header Set/Get is invalid %s(expected: %s)", retval, testHeader)
	}

	header.Del("x-MBOX-HeadeR")

	retval = header.Get("x-mbox-HeadeR")
	if retval != "" {
		t.Fatalf("Header Set/Get is invalid %s(expected: %s)", retval, "")
	}
}
