package lsd

func Sort(a []string, w int) {
	n := len(a)
	r := 256
	aux := make([]string, n)

	for d := w - 1; d >= 0; d-- {
		count := make([]int, r+1)
		for i := 0; i < n; i++ {
			count[a[i][d]+1]++
		}

		for i := 0; i < r; i++ {
			count[i+1] += count[i]
		}

		for i := 0; i < n; i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}

		for i := 0; i < n; i++ {
			a[i] = aux[i]
		}
	}
}
