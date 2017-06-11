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

import "testing"

func TestSeedRandom(t *testing.T) {
	alpha := NewCryptoSeededRandom()
	omega := NewCryptoSeededRandom()

	for _, prime := range primes {
		a := alpha.Intn(prime)
		o := omega.Intn(prime)

		// Yes these could end up being equal, but it will
		// only happen very infrequently.
		// https://en.wikipedia.org/wiki/Birthday_problem
		if a == o {
			t.Errorf("unexpected %d == %d", a, o)
		}

		// Even SMALLER chance of collision
		a64 := alpha.Uint64()
		o64 := omega.Uint64()
		if a64 == o64 {
			t.Errorf("unexpected %d == %d", a64, o64)
		}
	}
}
