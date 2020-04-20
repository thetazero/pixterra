package main

import (
	"math/rand"
)

func generateChunk(x, y int) chunk {
	c := chunk{}
	seed := int64(pRandom(x, y) * 100000)
	// fmt.Println(seed)
	rand.Seed(seed) //bad
	for y := 0; y < 8; y++ {
		c = append(c, []int{})
		for x := 0; x < 8; x++ {
			r := rand.Intn(11)
			n := 0
			switch {
			case r == 9:
				n = 2
			case r > 8:
				n = 3
			case r > 2:
				n = 1
			default:
				n = 0
			}
			c[y] = append(c[y], n)
		}
	}
	return c
}
