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
)

// NewCryptoSeededSource builds a "math/rand" Source that uses "crypto/rand"
// to generate a non-deterministic seed. Note: this Source still outputs
// a deterministic sequence based on the seed, it's just that the seed is
// obfuscated.
func NewCryptoSeededSource() rand.Source {
	var value int64
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return rand.NewSource(value)
}

// NewCryptoSeededRandom is a convenience builder around NewCryptoSeededSource()
// that returns a pointer to a "math/rand" Rand that is ready to use.
func NewCryptoSeededRandom() *rand.Rand {
	return rand.New(NewCryptoSeededSource())
}
