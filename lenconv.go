package lenconv

import (
	"fmt"
)

type Ft float64
type Met float64

func (ft Ft) String() string {
	return fmt.Sprintf("%g Ft.\n", ft)
}

func (m Met) String() string {
	return fmt.Sprintf("%g m\n", m)
}

// Переводим Метры в Футы
// Делим длину на 3.281
func FtToMet(ft Ft) Met {
	return Met(ft / 3.281)
}

// Переводим Футы в Метры
// Умножаем длину на 3.281
func MetToFt(m Met) Ft {
	return Ft(m * 3.281)
}
