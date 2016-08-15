// Copyright © 2016 shoarai

package archive

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// ErrFormat indicates that decoding encountered an unknown format.
var ErrFormat = errors.New("archive: unknown format")

type format struct {
	name, magic string
	read        func(file *os.File, out io.Writer) error
}

var formats = []format{}

// A reader is an io.Reader that can also peek ahead.
type reader interface {
	io.Reader
	Peek(int) ([]byte, error)
}

// Formats is the list of registered formats.
func RegistorFormat(name, magic string, read func(file *os.File, out io.Writer) error) {
	formats = append(formats, format{name, magic, read})
}

func Read(file *os.File, out io.Writer) error {
	r := bufio.NewReader(file)
	rr := asReader(r)
	f := sniff(rr)
	if f.read == nil {
		return ErrFormat
	}
	err := f.read(file, out)
	return err
}

// asReader converts an io.Reader to a reader.
func asReader(r io.Reader) reader {
	if rr, ok := r.(reader); ok {
		return rr
	}
	return bufio.NewReader(r)
}

// Match reports whether magic matches b. Magic may contain "?" wildcards.
func match(magic string, b []byte) bool {
	if len(magic) != len(b) {
		return false
	}
	for i, c := range b {
		if magic[i] != c && magic[i] != '?' {
			return false
		}
	}
	return true
}

// Sniff determines the format of r's data.
func sniff(r reader) format {
	for _, f := range formats {
		b, err := r.Peek(len(f.magic))
		if err == nil && match(f.magic, b) {
			return f
		}
	}
	return format{}
}
