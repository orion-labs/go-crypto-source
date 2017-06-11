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
	mrand "math/rand"
	"sync"
)

type cryptSrc struct {
	sync.Mutex
	safe bool
	buf  []byte
}

func (s *cryptSrc) Seed(seed int64) { /*no-op*/ }

func (s *cryptSrc) Int63() int64 {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	crand.Read(s.buf)
	s.buf[0] &= 0x7f
	return int64(binary.BigEndian.Uint64(s.buf))
}

func (s *cryptSrc) Uint64() uint64 {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}
	crand.Read(s.buf)
	return binary.BigEndian.Uint64(s.buf)
}

// NewSource builds struct that conforms to the `math/rand` `Source64` interface,
// and provides a non-deterministic random numbers as provided by `crypto/rand`.
// This is set up to have minimal allocations by sharing a single buffer, so
// you are required to specify whether or not you want thread safety.
func NewSource(threadsafe bool) mrand.Source64 {
	return &cryptSrc{safe: threadsafe, buf: make([]byte, 8)}
}

// NewRandom is a convenience builder around `NewSource(...)` that returns a
// `math/rand` `*Rand` struct that is directly ready for use.
func NewRandom(threadsafe bool) *mrand.Rand {
	return mrand.New(NewSource(threadsafe))
}
