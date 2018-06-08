package fork

import (
	"bytes"
	"fmt"
	"io"
)

// Reader forks the src Reader a given number of times.
func Reader(src io.Reader, number int) []io.Reader {
	var cache []byte
	var buff bytes.Buffer
	_, err := buff.ReadFrom(src)
	if err != nil {
		fmt.Println("error found", err)
		return nil
	}
	cache = buff.Bytes()
	readers := make([]io.Reader, 10)
	for i := range readers {
		readers[i] = &forkReader{0, cache}
	}
	return readers
}

type forkReader struct {
	readPosition int
	cache        []byte
}

func (fr *forkReader) Read(p []byte) (n int, err error) {
	need := len(p)
	avail := len(fr.cache) - fr.readPosition
	if avail < need {
		n = avail
		err = io.EOF
	} else {
		n = need
		err = nil
	}
	for i := fr.readPosition; i < fr.readPosition+n; i++ {
		p[i-fr.readPosition] = fr.cache[i]
	}
	fr.readPosition += n
	return n, err
}
