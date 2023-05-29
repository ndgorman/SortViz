package main

func validateSort(pSlc *[]float64) bool {
	slc := *pSlc
	for i := 0; i < len(slc)-1; i++ {
		if slc[i] > slc[i+1] {
			return false
		}
	}
	return true
}

func insertSort(pSlc *[]float64) {
	slc := *pSlc
	for i := 1; i < len(slc); i++ {
		for j := i; j > 0 && slc[j] < slc[j-1]; j-- {
			slc[j-1], slc[j] = slc[j], slc[j-1]
		}
	}

}

func quickSort(pSl *[]float64, p int, r int, output chan []float64) {
	a := *pSl
	if p < r {
		//partition
		q := func() int {
			i := p - 1
			x := a[r]
			for j := p; j < r; j++ {
				if a[j] <= x {
					i++
					a[i], a[j] = a[j], a[i]
					output <- a
				}
			}
			a[i+1], a[r] = a[r], a[i+1]
			output <- a
			return i + 1
		}()
		quickSort(pSl, p, q-1, output)
		quickSort(pSl, q+1, r, output)
	}
}
