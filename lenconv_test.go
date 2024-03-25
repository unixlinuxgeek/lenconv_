package lenconv

import (
	"math/rand"
	"testing"
)

// Модульный тест проверяет работу метода FtToMet
func TestFtToMet(t *testing.T) {
	r := float64(rand.Intn(10))
	res := r / 3.281

	t.Logf("Convert random: %g Ft. to Meter == %g Meter\n", r, float64(FtToMet(Ft(r))))

	if res == float64(FtToMet(Ft(r))) {
		t.Log("TestFtToMet Passed!!!")
	} else {
		t.Error("TestFtToMet Failed")
	}
}

// Модульный тест проверяет работу метода MetToFt
func TestMetToFt(t *testing.T) {
	r := float64(rand.Intn(10))
	res := r * 3.281

	t.Logf("Convert random: %g Met to Ft. == %g Ft.\n", r, float64(MetToFt(Met(r))))

	if res == float64(MetToFt(Met(r))) {
		t.Log("TestMetToFt Passed!!!")
	} else {
		t.Error("TestMetToFt Failed")
	}
}
