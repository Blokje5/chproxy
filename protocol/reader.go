package protocol

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type Reader struct {
	buf []byte
	n int
	maxSize int
}

func NewReader(buf []byte) *Reader {
	return &Reader{
		buf: buf,
		maxSize: len(buf),
	}
}

func (r *Reader) ensureCanReadN(n int) error {
	if n > (r.maxSize - r.n) {
		return ErrUnexpectedEOF
	}

	return nil
}


func (r *Reader) ReadUint32() (uint32, error) {
	err := r.ensureCanReadN(Uint32Size)
	if err != nil {
		return 0, err
	}

	v := binary.LittleEndian.Uint32(r.buf[r.n:])
	r.n += Uint32Size

	return v, nil
}

func (r *Reader) ReadUint64() (uint64, error) {
	err := r.ensureCanReadN(Uint32Size)
	if err != nil {
		return 0, err
	}
	
	v := binary.LittleEndian.Uint64(r.buf[r.n:])
	r.n += 8

	return v, nil
}

func (r *Reader) ReadString() (string, error) {
	reader := bytes.NewReader(r.buf[r.n:])
	initial := reader.Len()
	size, err := binary.ReadUvarint(reader)
	if err != nil {
		return "", err
	}

	// calculate read bytes
	r.n += (initial - reader.Len())

	buf := make([]byte, size)

	n, err := io.ReadFull(reader, buf)
	if  err != nil {
		if errors.Is(err, io.EOF) {
			return "", ErrUnexpectedEOF
		}

		return "", err
	}
	r.n += n

	return string(buf), nil
}
