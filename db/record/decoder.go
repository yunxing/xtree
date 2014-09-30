package record

import (
	"encoding/binary"
	"errors"
	"io"
)

type RecordDecoder struct {
	r io.Reader
}

func (rd *RecordDecoder) Decode(rec *Record) error {
	var crc, length uint32

	// Read CRC
	err := binary.Read(rd.r, binary.LittleEndian, &crc)
	if err != nil {
		return err
	}

	// Read length
	err = binary.Read(rd.r, binary.LittleEndian, &length)
	if err != nil {
		return err
	}

	// Read data
	rec.data = make([]byte, length)
	_, err = io.ReadFull(rd.r, rec.data)
	if err != nil {
		return err
	}

	if ok := rec.vfyCRC(crc); !ok {
		return errors.New("crc unmatch, data corrupted")
	}

	return nil
}

func NewRecordDecoder(r io.Reader) *RecordDecoder {
	return &RecordDecoder{r}
}
