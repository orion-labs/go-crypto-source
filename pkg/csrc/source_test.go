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

var primes = []int{
	7573, 7577, 7583, 7589, 7591, 7603, 7607, 7621, 7639, 7643,
	7649, 7669, 7673, 7681, 7687, 7691, 7699, 7703, 7717, 7723,
	7727, 7741, 7753, 7757, 7759, 7789, 7793, 7817, 7823, 7829,
	7841, 7853, 7867, 7873, 7877, 7879, 7883, 7901, 7907, 7919,
}

func TestRandomSource(t *testing.T) {
	alpha := NewRandom()
	omega := NewRandom()

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
