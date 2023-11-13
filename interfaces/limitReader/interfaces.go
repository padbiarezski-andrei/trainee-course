package interfaces

import "io"

type myLimitReader struct {
	r io.Reader
	n int64
}

func (l *myLimitReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if l.n < int64(len(p)) {
		p = p[:l.n]
	}
	n, err = l.r.Read(p)
	l.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &myLimitReader{r, n}
}
