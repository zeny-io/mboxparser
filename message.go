package mboxparser

import (
	"bytes"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
)

type Message struct {
	Header Header
	Bodies []*Body
}

type Header map[string][]string

type Body struct {
	Header   Header
	FileName string
	Content  io.Reader
}

func (h Header) Get(key string) string {
	if s, ok := h[http.CanonicalHeaderKey(key)]; ok && len(s) > 0 {
		return s[0]
	}
	return ""
}

func (h Header) Set(key, val string) {
	h[http.CanonicalHeaderKey(key)] = []string{val}
}

func (h Header) Del(key string) {
	delete(h, http.CanonicalHeaderKey(key))
}

func newBodyByPart(part *multipart.Part) *Body {
	body := &Body{
		Header:   Header{},
		FileName: "",
		Content:  nil,
	}

	for k, v := range part.Header {
		body.Header[http.CanonicalHeaderKey(k)] = decodeHeader(v)
	}

	body.FileName = part.FileName()

	buf := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buf, part)
	if err != nil {
		return nil
	}

	mediaType, params, err := mime.ParseMediaType(body.Header.Get("Content-Type"))
	if err != nil {
		return nil
	}

	charset := params["charset"]
	encoding := body.Header.Get("Content-Transfer-Encoding")

	body.Header.Set("Content-Type", mediaType)

	body.Content = newDecoder(buf, charset, encoding)

	return body
}

func newBodyByMessage(message *Message, r io.Reader) *Body {
	body := &Body{
		Header:   Header{},
		FileName: "",
		Content:  nil,
	}

	for _, k := range []string{"Content-Type", "Content-Transfer-Encoding"} {
		body.Header.Set(k, message.Header.Get(k))
		message.Header.Del(k)
	}

	mediaType, params, err := mime.ParseMediaType(body.Header.Get("Content-Type"))
	if err != nil {
		return nil
	}

	charset := params["charset"]
	encoding := body.Header.Get("Content-Transfer-Encoding")

	body.Header.Set("Content-Type", mediaType)

	body.Content = newDecoder(r, charset, encoding)

	return body
}
