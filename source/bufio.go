
package bufio

import (

)

const (
	defaultBufSize = 4096
)

var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnusedByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

//Reader implements buffering for an io.Reader object
type Reader struct {
	buf 		  []byte
	rd 			  io.Reader  //reader provided by the client
	r,w 		  int //buf read and write positions
	err           error
	lastByte      int
	lastRuneSize  int
}


const (
	minReadBufferSize = 16
	maxConsecutiveEmptyReads = 100
)

func NewReaderSize(rd io.Reader, size int) *Reader {
	b, ok := rd.(*Reader)
	if ok && len(b.buf) >= size {
		return b
	}

	if size < minReadBufferSize {
		size = minReadBufferSize
	}

	r := new(Reader)
	r.reset(make([]byte, size), rd)
	return r
}


func (b *Reader) Reset(buf []byte, r io.Reader) {
	*b = Reader{
		buf : buf,
		rd : r,
		lastByte : -1,
		lastRuneSize : -1,
	}
}

func NewReader(rd io.Reader) *Reader{
	return NewReaderSize(rd, defaultBufSize)
}

func (b *Reader) reset(buf []byte, r io.Reader) {
	*b = Reader {
		buf : buf,
		rd : r,
		lastByte : -1,
		lastRuneSize : -1,
	}
}

var errNegativeRead = errors.New("bufio: reader returned negative count from Read")

func (b *Reader) fill() {
	if b.r > 0{
		copy(b.buf, b.buf[b.r: b.w])
		b.w -= b.r
		b.r = 0
	}

	if b.w >= len(b.buf) {
		panic("bufio: tried to fill full buffer")
	}

	for i := maxConsecutiveEmptyReads; i > 0; i--{
		n, err := b.rd.Read(b.buf[b.w:])
		if n < 0 {
			panic (errNegativeRead)
		}
		b.w += n
		if err != nil {
			b.err = err
			return
		}
		if n > 0 {
			return 
		}
	}
	b.err = io.ErrNoProgress
}


