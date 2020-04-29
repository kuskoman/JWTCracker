package generators

type WordGenerator struct {
	Alphabet string
	CurrChs  []int
}

func (w *WordGenerator) Next() string {
	i := len(w.CurrChs) - 1
	for i >= 0 {
		if i == 0 && w.CurrChs[i] == len(w.Alphabet)-1 {
			w.CurrChs[i] = 0
			w.CurrChs = append(w.CurrChs, 0)
			break
		}
		if w.CurrChs[i] < (len(w.Alphabet) - 1) {
			w.CurrChs[i] += 1
			break
		}
		w.CurrChs[i] = 0
		i--
	}

	var str string
	j := 0
	for j < len(w.CurrChs) {
		n := w.CurrChs[j]
		str += string(w.Alphabet[n])
		j++
	}

	return str
}

func NewAlphabeticNumerator() *WordGenerator {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var chs []int
	chs = append(chs, -1)
	w := WordGenerator{Alphabet: alphabet, CurrChs: chs}
	return &w
}
