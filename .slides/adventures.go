package main

import (
	crand "crypto/rand"
	"encoding/binary"
	mrand "math/rand"
)

/*
START REF03a OMIT
func New(src Source) *Rand
func NewSource(seed int64) Source
END REF03a OMIT

START REF03b OMIT
type ByteOrder interface {
	...
    Uint64([]byte) uint64
	...
}

// BigEndian is the big-endian implementation of ByteOrder.
var BigEndian bigEndian
END REF03b OMIT
*/
// START SAMPLE04 OMIT
func NewCryptoSeededSource() mrand.Source {
	b := make([]byte, 8)
	crand.Reader.Read(b)
	return mrand.NewSource(int64(binary.BigEndian.Uint64(b)))
}

// END SAMPLE04 OMIT

type MySource struct{}

// START MYSOURCESEED OMIT
func (m *MySource) Seed(int64) { /* no-op */ }

// END MYSOURCESEED OMIT

// START MYSOURCE63 OMIT
func (m *MySource) Int63() int64 {
	buf := make([]byte, 8)
	crand.Read(buf)
	buf[0] &= 0x7f
	return int64(binary.BigEndian.Uint64(buf))
}

// END MYSOURCE63 OMIT

// START MYSOURCE64 OMIT
func (m *MySource) Uint64() uint64 {
	buf := make([]byte, 8)
	crand.Read(buf)
	return binary.BigEndian.Uint64(buf)
}

// END MYSOURCE64 OMIT
