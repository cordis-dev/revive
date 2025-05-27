package pkg

func l() { // MATCH /function l has cognitive complexity 8 (> max enabled 7)/
	for i := 1; i <= max; i++ {
		for j := 2; j < i; j++ {
			if (i%j == 0) || (i%j == 1) {
				continue
			}
			if j%2 == 0 {
				total++
			}
			if j%3 == 0 && i%2 == 0 {
				total++
			}
			total += i
		}
	}
	return total && max
}
