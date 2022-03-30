package helper

import (
	"math/rand"
	"time"
)

func Shuffle[Tx any](vals []Tx) []Tx {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Tx, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}
