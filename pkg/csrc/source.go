// Copyright 2017 Orion Labs, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package csrc

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"sync"
)

const mask uint64 = ^uint64(1 << 63)

func to63(in uint64) int64 {
	return int64(in & mask)
}

type cryptSrc struct {
	sync.Mutex
	safe bool
	buf  []byte
}

func (s *cryptSrc) Seed(seed int64) { /*no-op*/ }

func (s *cryptSrc) Uint64() uint64 {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	crand.Read(s.buf)
	return binary.BigEndian.Uint64(s.buf)
}

func (s *cryptSrc) Int63() int64 {
	return to63(s.Uint64())
}

// NewSource builds a struct that conforms to the "math/rand" Source64 interface,
// and provides a non-deterministic random numbers as provided by "crypto/rand".
// This is set up to have minimal allocations by sharing a single buffer, so
// you are required to specify whether or not you want thread safety.
func NewSource(threadsafe bool) rand.Source64 {
	return &cryptSrc{safe: threadsafe, buf: make([]byte, 8)}
}

// NewRandom is a convenience builder around NewSource(...) that returns a
// pointer to a "math/rand" Rand instance that is directly ready for use.
func NewRandom(threadsafe bool) *rand.Rand {
	return rand.New(NewSource(threadsafe))
}

type simpleSrc struct{}

func (s *simpleSrc) Seed(seed int64) { /*no-op*/ }

func (s *simpleSrc) Uint64() (value uint64) {
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return value
}

func (s *simpleSrc) Int63() int64 {
	return to63(s.Uint64())
}

// NewSimpleSource builds a struct that conforms to the "math/rand" Source64
// interface, and provides a non-deterministic random numbers as provided by
// "crypto/rand".
//
// Calling out to the "crypto/rand" package is demonstrably slower than
// using the deterministically generated numbers from the "math/rand" package,
// so if performance is your intention, reconsider using this.
//
// If you only need a truly random seed for instantiating a "math/rand" Rand,
// then see NewCryptoSeededSource() or NewCryptoSeededRandom()
//
// If you are trying to reduce memory allocations (but are okay with paying the
// "crypto/rand" speed hit), see NewSource(...) or NewRandom(...).
func NewSimpleSource() rand.Source64 {
	return &simpleSrc{}
}

// NewSimpleRandom is a convenience builder around NewSimpleSource() that
// returns a pointer to a "math/rand" Rand instance that is directly ready
// for use.
func NewSimpleRandom() *rand.Rand {
	return rand.New(NewSimpleSource())
}
