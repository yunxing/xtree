package record

import (
	"bytes"
	"io"
)

type Reader struct {
	rs io.ReadSeeker
}

func NewReader(rs io.ReadSeeker) *Reader {
	return &Reader{rs: rs}
}

func (rd *Reader) Read(index int64) (io.Reader, error) {
	_, err := rd.rs.Seek(index, 0)
	if err != nil {
		return nil, err
	}

	rec := Record{}
	dec := NewRecordDecoder(rd.rs)
	err = dec.Decode(&rec)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(rec.data), nil
}
