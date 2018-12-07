package rand_helper

import (
	"math/rand"
	"time"
)

var (
	Uint64 = rand.Uint64
	Int63  = rand.Int63
	Int31  = rand.Int31
	Intn   = rand.Intn
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
