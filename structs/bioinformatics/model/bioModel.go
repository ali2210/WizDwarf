package model

type Levenshtein struct {
	Probablity float32
	Percentage float32
	Name       string
}

type LevenshteinInterface interface {
	SetProbParameter(p float32)
	GetProbParameter() float32
	Result(d int) float32
	CalcualtePercentage(p float32) float32
}

func (l *Levenshtein) SetProbParameter(p float32) {
	(*l).Probablity = p
}

func (l *Levenshtein) GetProbParameter() float32 {
	if (*l).Probablity >= 0.0 {
		return (*l).Probablity
	}
	return -1.0
}
func (*Levenshtein) Result(d int) float32 {
	return float32(d / 1000)
}

func (*Levenshtein) CalcualtePercentage(p float32) float32 {
	return (((100 - p) / 100) * 100)
}
