package protocol

import (
	"bytes"
	"encoding/binary"
)

type Writer struct {
	buf []byte
	n int
	maxSize int
}
	
func NewWriter(buf []byte) *Writer {
	return &Writer{
		buf: buf,
		maxSize: len(buf),
	}
}

func (w *Writer) PutByte(v byte) error {
	w.buf[w.n] = v
	w.n += 1

	return nil
}

func (w *Writer) PutUint32(v uint32) error {
	binary.LittleEndian.PutUint32(w.buf[w.n:], v)
	w.n += Uint32Size

	return nil
}

func (w *Writer) PutUint64(v uint64) error {
	binary.LittleEndian.PutUint64(w.buf[w.n:], v)
	w.n += Uint64Size

	return nil
}

func (w *Writer) PutString(v string) error {
	n := binary.PutUvarint(w.buf, uint64(len(v)))
	w.n += n

	buffer := bytes.NewBuffer(w.buf)
	n, err := buffer.WriteString(v)
	w.n += n

	return err
}
