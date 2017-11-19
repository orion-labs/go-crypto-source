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
	"math/rand"
	"testing"
	"time"
)

var prime1000 = 7919

var result int

func BenchmarkGlobal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = rand.Intn(prime1000)
	}
}

func BenchmarkNative(b *testing.B) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for n := 0; n < b.N; n++ {
		result = random.Intn(prime1000)
	}
}

func BenchmarkSeed(b *testing.B) {
	random := NewCryptoSeededRandom()
	for n := 0; n < b.N; n++ {
		result = random.Intn(prime1000)
	}
}

func BenchmarkCryptoRead(b *testing.B) {
	buffer := make([]byte, 8)
	for n := 0; n < b.N; n++ {
		result, _ = crand.Read(buffer)
	}
}

func BenchmarkCryptoUnsafe(b *testing.B) {
	random := NewRandom(false)
	for n := 0; n < b.N; n++ {
		result = random.Intn(prime1000)
	}
}

func BenchmarkCryptoSafe(b *testing.B) {
	random := NewRandom(true)
	for n := 0; n < b.N; n++ {
		result = random.Intn(prime1000)
	}
}

func BenchmarkCryptoSimple(b *testing.B) {
	random := NewSimpleRandom()
	for n := 0; n < b.N; n++ {
		result = random.Intn(prime1000)
	}
}
