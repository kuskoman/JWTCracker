package main

type WordGenerator struct {
	alphabet string
	currChs  []int
}

func (w *WordGenerator) Next() string {
	i := len(w.currChs) - 1
	for i >= 0 {
		if i == 0 && w.currChs[i] == len(w.alphabet)-1 {
			w.currChs[i] = 0
			w.currChs = append(w.currChs, 0)
			break
		}
		if w.currChs[i] < (len(w.alphabet) - 1) {
			w.currChs[i] += 1
			break
		}
		w.currChs[i] = 0
		i--
	}

	var str string
	j := 0
	for j < len(w.currChs) {
		n := w.currChs[j]
		str += string(w.alphabet[n])
		j++
	}

	return str
}
