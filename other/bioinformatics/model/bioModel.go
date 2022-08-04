package model

import "reflect"

type Levenshtein struct {
	Probablity float32
	Percentage float32
	Name       string
}

type LevenshteinInterface interface {
	SetProbParameter(p float32)
	GetProbParameter() float32
	Result(outcome, total_space int) float32
	CalcualtePercentage(value float32) float32
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
func (l *Levenshtein) Result(outcome, total_space int) float32 {

	if reflect.DeepEqual(outcome, 0) {
		return -1.00
	}

	l.SetProbParameter(float32(outcome / total_space))
	return l.GetProbParameter()
}

func (l *Levenshtein) CalcualtePercentage(value float32) float32 {

	if reflect.DeepEqual(value, 0) {
		return -1.00
	}

	if reflect.DeepEqual(value, 100) {
		return 0.00
	}

	return value * 10
}
