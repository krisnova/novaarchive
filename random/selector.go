package random

import (
	"crypto"
	"fmt"
	"hash/fnv"
	"math"
	"math/rand"
	"os/exec"
	"runtime"
	"time"
)

// Selector is a way to generate random items for a set
// of anything
type Selector struct {
	v []interface{}
}

// NewSelector is a memory and system intensive
// struct that can be used to generate random items
// from a set
func NewSelector(v []interface{}) *Selector {
	return &Selector{
		v: v,
	}
}

// Pick is used to pick a random item from an initialized
// Selctor.
func (s *Selector) Pick() interface{} {
	if len(s.v) < 1 {
		return nil
	}
	if len(s.v) == 1 {
		return s.v[0]
	}
	r := IBetween(0, len(s.v))
	return s.v[r]
	return nil
}

// IBetween uses system entropy to seed
// the standard library rand package
// and pick a random value between X, and N
func IBetween(x, n int) int {
	if x == n {
		return 0
	}
	d := int(math.Abs(float64(x - n)))
	rand.Seed(NovaRandomInt64())
	r := rand.Intn(d)
	if x > n {
		return n + r
	}
	return x + r
}

// NovaRandomInt64 will wrap up system commands
// and calculate hash sums for a pathetic attempt
// at generating entropy
func NovaRandomInt64() int64 {
	if runtime.GOOS == "linux" {
		// Thanks Cesar
		// openssl rand -base64 32
		rBytes, _ := exec.Command("openssl", "rand", "-base64", "32").Output()
		hashSSL := fnv.New32a()
		hashSSL.Write(rBytes)
		sumSSL := int64(hashSSL.Sum32())
		// Thanks Stephanie
		procBytes, _ := exec.Command("ls", "-lR", "/proc", "").Output()
		hashProc := fnv.New32a()
		hashProc.Write(procBytes)
		sumProc := int64(hashSSL.Sum32())
		return hashint64(sumProc, sumSSL, time.Now().UnixNano())
	}
	return time.Now().UnixNano()
}

// hashint64 will take a set of int64 values
// and MD5 sum them together and calculate
// a int64 value of the sum
func hashint64(iset ...int64) int64 {
	h := crypto.MD5.New()
	for _, i64 := range iset {
		fmt.Fprintf(h, "%d", i64)
	}
	iHash := fnv.New32a()
	iHash.Write(h.Sum(nil))
	return int64(iHash.Sum32())
}
