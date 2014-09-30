package record

import (
	"encoding/binary"
	"hash/crc32"
	"unsafe"
)

const (
	sizeOfLength = int(unsafe.Sizeof(uint32(0)))
	sizeOfCRC    = int(unsafe.Sizeof(uint32(0)))
)

var crcTable = crc32.MakeTable(crc32.Koopman)

type Record struct {
	data []byte
}

func (r *Record) calcCRC() uint32 {
	size := len(r.data)
	lb := make([]byte, sizeOfLength)
	binary.LittleEndian.PutUint32(lb, uint32(size))
	crc := crc32.Checksum(lb, crcTable)
	crc = crc32.Update(crc, crcTable, r.data)
	return crc
}

func (r *Record) vfyCRC(crc uint32) bool {
	if crc != r.calcCRC() {
		return false
	}
	return true
}
