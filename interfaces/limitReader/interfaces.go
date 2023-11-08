package interfaces

import "io"

/*
Функция LimitReader из пакета io принимает переменную r типа io .Reader и количество байтов n и возвращает другой объект Reader, который читает из r, но после чтения n сообщает о достижении конца файла.
Реализуйте его.
func LimitReader(r io.Reader, n int64) io.Reader
Для проверки используйте тип bytes.Buffer, который реализует интерфейсы io.Writer, io.Reader
*/

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
