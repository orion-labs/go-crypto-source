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
)

type cryptSrc struct{}

func (s *cryptSrc) Seed(seed int64) { /*no-op*/ }

func (s *cryptSrc) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *cryptSrc) Uint64() (value uint64) {
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return value
}

// NewSource builds struct that conforms to the `math/rand` `Source64` interface,
// and provides a non-deterministic random numbers as provided by `crypto/rand`.
func NewSource() mrand.Source64 {
	return &cryptSrc{}
}

// NewRandom is a convenience builder around `NewSource(...)` that returns a
// `math/rand` `*Rand` struct that is directly ready for use.
func NewRandom() *mrand.Rand {
	return mrand.New(NewSource())
}
