package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
type Personal struct {
	Name   string
	Weight float64 // вес в килограммах.
	Height float64 // рост в метрах.
}

// Print выводит персональные данные пользователя.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
