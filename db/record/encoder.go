package record

import (
	"encoding/binary"
	"io"
)

type RecordEncoder struct {
	w io.Writer
}

func (re *RecordEncoder) Encode(rec *Record) error {
	// Write CRC
	if _, err := re.w.Write(encCRC(rec)); err != nil {
		return err
	}
	// Write length
	if _, err := re.w.Write(encLength(rec)); err != nil {
		return err
	}
	// Write data
	if _, err := re.w.Write(rec.data); err != nil {
		return err
	}
	return nil
}

func NewRecordEncoder(w io.Writer) *RecordEncoder {
	return &RecordEncoder{w}
}

func encLength(rec *Record) []byte {
	size := len(rec.data)
	lb := make([]byte, sizeOfLength)
	binary.LittleEndian.PutUint32(lb, uint32(size))
	return lb
}

func encCRC(rec *Record) []byte {
	crcb := make([]byte, sizeOfCRC)
	binary.LittleEndian.PutUint32(crcb, rec.calcCRC())
	return crcb
}
