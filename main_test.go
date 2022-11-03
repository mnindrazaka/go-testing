package main

import "testing"

var (
	kubusKu            Kubus   = NewKubus(4)
	volumeSeharusnya   float64 = 64
	luasSeharusnya     float64 = 96
	kelilingSeharusnya float64 = 48
)

func TestHitungVolume(t *testing.T) {
	t.Logf("Volume : %.2f", kubusKu.Volume())
	if kubusKu.Volume() != volumeSeharusnya {
		t.Errorf("Salah! harusnya %.2f", volumeSeharusnya)
	}
}

func TestHitungLuas(t *testing.T) {
	t.Logf("Luas : %.2f", kubusKu.Luas())
	if kubusKu.Luas() != luasSeharusnya {
		t.Errorf("Salah! harusnya %.2f", luasSeharusnya)
	}
}

func TestHitungKeliling(t *testing.T) {
	t.Logf("Keliling : %.2f", kubusKu.Keliling())
	if kubusKu.Keliling() != kelilingSeharusnya {
		t.Errorf("Salah! harusnya %.2f", kelilingSeharusnya)
	}
}
