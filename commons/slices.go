package commons

func Map[t any](a []t, f func(t) t) []t {
	for i, entry := range a {
		a[i] = f(entry)
	}

	return a
}

func Reduce[t comparable](a []t, f func(t, t) t) t {
	var result t
	for _, v := range a {
		// if i == 0 {
		// 	continue
		// }

		result = f(result, v)
	}

	return result
}
