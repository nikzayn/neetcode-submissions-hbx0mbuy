func shipWithinDays(weights []int, days int) int {
	l, r := 0, 0

	for _, w := range weights {
		if w > l {
			l = w
		}
		r += w
	}

	res := r

	for l <= r {
		k := (l + r) / 2

		if canShip(weights, days, k){
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return res
}

func canShip(weights []int, days, capacity int) bool {
	daysNeeded := 1
	curr := 0

	for _, w := range weights {
		if curr+w > capacity {
			daysNeeded++
			curr = 0
		}
		curr += w
	}
	return daysNeeded <= days
}
