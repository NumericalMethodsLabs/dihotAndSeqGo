package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Sin(x)
}

//Метод секущей
func sek() {
	var a, b, eps float64
	fmt.Println("Введите границы отрезка:")
	fmt.Scan(&a, &b)
	a, b = math.Min(a, b), math.Max(a, b)

	fmt.Println("Введите погрешность:")
	fmt.Scan(&eps)

	c1 := (a*f(b) - b*f(a)) / (f(b) - f(a))
	c2 := c1 * 2
	if f(a)*f(b) > 0 {
		fmt.Println("Корней нет")
	} else {
		count := 0
		for math.Abs(f(c1)-f(c2)) >= eps {
			if f(a)*f(c1) < 0 {
				b = c1
			} else {
				a = c1
			}
			c2 = c1
			c1 = (a*f(b) - b*f(a)) / (f(b) - f(a))
			count += 1
		}

		fmt.Printf("Результат: %f \n", c1)
		fmt.Printf("Отклонение: %f \n", math.Abs(math.Pi-c1))
		fmt.Printf("Количество итераций: %d \n", count)
	}
}

//Дихотомия
func dikhotomiya() {
	var a, b, eps float64
	fmt.Println("Введите границы отрезка:")
	fmt.Scan(&a, &b)
	a, b = math.Min(a, b), math.Max(a, b)

	fmt.Println("Введите погрешность:")
	fmt.Scan(&eps)

	c := (a + b) / 2
	if f(a)*f(b) > 0 {
		fmt.Println("Корней нет")
	} else {
		count := 0
		for math.Abs(b-a) > eps {
			if f(a)*f(c) < 0 {
				b = c
			} else {
				a = c
			}
			c = (a + b) / 2
			count += 1
		}

		fmt.Printf("Результат: %f \n", c)
		fmt.Printf("Отклонение: %f \n", math.Abs(math.Pi-c))
		fmt.Printf("Количество итераций: %d \n", count)
	}
}

func main() {
	fmt.Println("Метод секущей:")
	sek()
	fmt.Println("Дихотомия:")
	dikhotomiya()
}
