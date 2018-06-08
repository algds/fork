package fork

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestForkReader(t *testing.T) {
	src := strings.NewReader("this is the string data source")
	forks := Reader(src, 2)
	x := make([]byte, 11)
	forks[0].Read(x)
	y := make([]byte, 11)
	forks[1].Read(y)
	if !bytes.Equal(x, y) {
		t.Errorf("data should be equal (%s) and (%s)", x, y)
	}
	x = make([]byte, 7)
	forks[0].Read(x)
	if !bytes.Equal(x, []byte(" string")) {
		t.Errorf("data should be equal (%s) and (%s)", x, " string")
	}
	x = make([]byte, 13)
	read, err := forks[0].Read(x)
	if read != 12 {
		t.Errorf("read should be %d but got %d", 12, read)
	}
	if err != io.EOF {
		t.Errorf("error should be %v but got %v", io.EOF, err)
	}
}

type badReader struct{}

func (*badReader) Read(p []byte) (n int, err error) {
	return 0, bytes.ErrTooLarge
}

func TestBadSourceReader(t *testing.T) {
	var br *badReader
	forks := Reader(br, 2)
	if forks != nil {
		t.Errorf("Forks should be nil but got %v", forks)
	}
}
