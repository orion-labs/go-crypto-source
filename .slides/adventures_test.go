package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
	"os"
	"testing"
	"time"
)

func TestExample01(t *testing.T) {
	// START EXAMPLE01 OMIT
	seed := int64(7)
	a := mrand.New(mrand.NewSource(seed))
	b := mrand.New(mrand.NewSource(seed))

	primes := []int{29, 31, 37, 41, 43, 47, 53, 59}
	for _, prime := range primes {
		fmt.Fprintf(os.Stdout, "%2d == %2d\n", a.Intn(prime), b.Intn(prime))
	}
	// END EXAMPLE01 OMIT

	/*
		START OUTPUT01 OMIT
		16 == 16
		20 == 20
		 7 ==  7
		19 == 19
		24 == 24
		 4 ==  4
		15 == 15
		52 == 52
		END OUTPUT01 OMIT
	*/
}

func TestExample02(t *testing.T) {
	// START EXAMPLE02 OMIT
	now := time.Now().UnixNano()
	c := mrand.New(mrand.NewSource(now))
	d := mrand.New(mrand.NewSource(now))
	time.Sleep(time.Millisecond)
	e := mrand.New(mrand.NewSource(time.Now().UnixNano()))

	primes := []int{29, 31, 37, 41, 43, 47, 53, 59}
	for _, prime := range primes {
		fmt.Fprintf(os.Stdout, "%2d == %2d != %2d\n", c.Intn(prime), d.Intn(prime), e.Intn(prime))
	}
	// END EXAMPLE02 OMIT

	/*
		START OUTPUT02 OMIT
		21 == 21 != 19
		 1 ==  1 != 27
		28 == 28 != 35
		18 == 18 != 27
		12 == 12 != 13
		17 == 17 != 41
		20 == 20 !=  1
		40 == 40 != 16
		END OUTPUT02 OMIT
	*/
}

func TestExampleBigEndian(t *testing.T) {
	// START BigEndEx OMIT
	buf := make([]byte, 8)
	for ix := 0; ix < 8; ix++ {
		crand.Read(buf)
		val := int64(binary.BigEndian.Uint64(buf))
		fmt.Fprintf(os.Stdout, "%x == %d\n", buf, val)
	}
	// END BigEndEx OMIT

	/*
		START BigEndOut OMIT
		857fcb4257ade141 == -8827113258823589567
		fcaaa0c59958c983 == -240202859569165949
		551d82c4fcb34fb8 == 6133202050113294264
		e250c4e3181204c8 == -2138993343360531256
		d18e97b0d21941e4 == -3346570687394790940
		0f40173c91f4efdc == 1098903857992626140
		b3cb649c847292c0 == -5491184697248410944
		0c904370fa8b03b2 == 905297677620282290
		END BigEndOut OMIT
	*/
}

func TestExampleSeeded(t *testing.T) {

	// START EXAMPLE04 OMIT
	f := mrand.New(NewCryptoSeededSource())
	g := mrand.New(NewCryptoSeededSource())

	primes := []int{29, 31, 37, 41, 43, 47, 53, 59}
	for _, prime := range primes {
		fmt.Fprintf(os.Stdout, "%2d != %2d\n", f.Intn(prime), g.Intn(prime))
	}
	// END EXAMPLE04 OMIT

	/*
		START OUTPUT04 OMIT
		 9 !=  2
		23 != 15
		28 != 23
		 0 != 23
		 4 != 12
		18 != 36
		28 != 17
		 9 != 35
		END OUTPUT04 OMIT
	*/
}

func TestExampleMySource(t *testing.T) {
	// START MYEXAMPLE OMIT
	h := mrand.New(&MySource{})
	i := mrand.New(&MySource{})

	primes := []int{29, 31, 37, 41, 43, 47, 53, 59}
	for _, prime := range primes {
		fmt.Fprintf(os.Stdout, "%2d != %2d\n", h.Intn(prime), i.Intn(prime))
	}
	// END MYEXAMPLE OMIT

	/*
		START MYOUTPUT OMIT
		20 != 24
		 4 !=  8
		 6 !=  9
		 5 !=  8
		 6 != 30
		 1 != 39
		39 !=  0
		10 != 49
		END MYOUTPUT OMIT
	*/

}
