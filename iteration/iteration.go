package iteration

func Repeat(s string, n int) (ns string) {
	for i := 0; i < n; i++ {
		ns += s
	}

	return
}
