package main

import "math/rand"

//return determanistic random number from x,y
func pRandom(x, y int) float64 {
	rand.Seed(int64(x))
	rand.Seed(int64(y)*rand.Int63() + int64(x))
	return rand.Float64()
}
