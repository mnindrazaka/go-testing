package main

import "math"

type kubus struct {
	Sisi float64
}

type Kubus interface {
	Volume() float64
	Luas() float64
	Keliling() float64
}

func NewKubus(sisi float64) Kubus {
	return kubus{Sisi: sisi}
}

func (kubus kubus) Volume() float64 {
	return math.Pow(kubus.Sisi, 3)
}

func (kubus kubus) Luas() float64 {
	return math.Pow(kubus.Sisi, 2) * 6
}

func (kubus kubus) Keliling() float64 {
	return kubus.Sisi * 12
}
