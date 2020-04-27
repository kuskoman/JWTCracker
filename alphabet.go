package main

type WordGenerator struct {
	alphabet string
	currChs  []int
}

func (w *WordGenerator) Next() string {
	i := len(w.currChs) - 1
	for i >= 0 {
		if i == 0 && w.currChs[i] == len(w.alphabet)-1 {
			w.currChs = append(w.currChs, 0)
			break
		}
		if w.currChs[i] < (len(w.alphabet) - 1) {
			w.currChs[i] += 1
			break
		}
		i--
	}

	var str string
	for chint := range w.currChs {
		str += string(w.alphabet[chint])
	}

	return str
}
