package main

func generateChunk(x, y int) chunk {
	c := chunk{}
	// seed := int64(pRandom(x, y) * 100000)
	// // fmt.Println(seed)
	// rand.Seed(seed) //bad
	xoff := x * 8
	yoff := y * 8
	for y := 0; y < 8; y++ {
		c = append(c, []int{})
		for x := 0; x < 8; x++ {
			// r := rand.Intn(11)
			r := simplex.Eval2(float64(x+xoff)/8, float64(yoff+y)/8)
			n := 0
			switch {
			case r > 0.9:
				n = 2
			case r > 0.7:
				n = 3
			case r > 0.3:
				n = 1
			default:
				n = 0
			}
			c[y] = append(c[y], n)
		}
	}
	return c
}
