package commons

func Map[i any, o any](a []i, f func(i) o) []o {
	b := make([]o, len(a))
	for i, entry := range a {
		b[i] = f(entry)
	}

	return b
}

func Reduce[t comparable](a []t, f func(t, t) t) t {
	var result t
	for _, v := range a {
		result = f(result, v)
	}

	return result
}
