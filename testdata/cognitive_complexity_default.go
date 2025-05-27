package pkg

func l() { // MATCH /function l has cognitive complexity 8 (> max enabled 7)/
	for i := 1; i <= max; i++ {
		for j := 2; j < i; j++ {
			if (i%j == 0) || (i%j == 1) {
				continue
			}
			total += i
		}
	}
	if total > 0 { // +1
		total--
	}
	if total > 0 && max > 0 { // +2
		total++
	}
	return total && max
}
