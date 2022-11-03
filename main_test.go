package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubusKu            Kubus   = NewKubus(4)
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	// t.Logf("Volume : %.2f", kubusKu.Volume())
	// if kubusKu.Volume() != volumeSeharusnya {
	// 	t.Errorf("Salah! harusnya %.2f", volumeSeharusnya)
	// }
	assert.Equal(t, kubusKu.Volume(), volumeSeharusnya, "perhitungan volume salah")
}

func TestHitungLuas(t *testing.T) {
	// t.Logf("Luas : %.2f", kubusKu.Luas())
	// if kubusKu.Luas() != luasSeharusnya {
	// 	t.Errorf("Salah! harusnya %.2f", luasSeharusnya)
	// }
	assert.Equal(t, kubusKu.Luas(), luasSeharusnya, "perhitungan luas salah")
}

func TestHitungKeliling(t *testing.T) {
	// t.Logf("Keliling : %.2f", kubusKu.Keliling())
	// if kubusKu.Keliling() != kelilingSeharusnya {
	// 	t.Errorf("Salah! harusnya %.2f", kelilingSeharusnya)
	// }
	assert.Equal(t, kubusKu.Keliling(), kelilingSeharusnya, "perhitungan keliling salah")
}
