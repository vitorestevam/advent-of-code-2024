package commons

func Map[t any](a []t, f func(t) t) []t {
	for i, entry := range a {
		a[i] = f(entry)
	}

	return a
}
