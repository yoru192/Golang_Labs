package main

import (
	"Lab_03/calc"
	"fmt"
)

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func (c Calc) Sum(nums ...float64) float64 {
	return calc.Sum(nums...)
}

func (c Calc) Max(nums ...float64) float64 {
	return calc.Max(nums...)
}

func (c Calc) Min(nums ...float64) float64 {
	return calc.Min(nums...)
}

func (c Calc) Divide(a, b float64) (float64, error) {
	return calc.Divide(a, b)
}

func task1() {
	fmt.Println("\n===== Завдання 1: пряме використання пакету calc =====")

	nums := []float64{3, 7, 1, 9, 4}

	fmt.Printf("Числа: %v\n", nums)
	fmt.Printf("Sum  : %.2f\n", calc.Sum(nums...))
	fmt.Printf("Max  : %.2f\n", calc.Max(nums...))
	fmt.Printf("Min  : %.2f\n", calc.Min(nums...))

	if result, err := calc.Divide(10, 4); err != nil {
		fmt.Println("Помилка:", err)
	} else {
		fmt.Printf("10 / 4 = %.2f\n", result)
	}

	if _, err := calc.Divide(5, 0); err != nil {
		fmt.Println("Очікувана помилка:", err)
	}
}

func task2() {
	fmt.Println("\n===== Завдання 2: використання через інтерфейс Calculator =====")

	var c Calculator = Calc{}

	nums := []float64{5.5, 2.1, 8.8, 3.3}

	fmt.Printf("Числа: %v\n", nums)
	fmt.Printf("Sum  : %.2f\n", c.Sum(nums...))
	fmt.Printf("Max  : %.2f\n", c.Max(nums...))
	fmt.Printf("Min  : %.2f\n", c.Min(nums...))

	if result, err := c.Divide(22, 7); err != nil {
		fmt.Println("Помилка:", err)
	} else {
		fmt.Printf("22 / 7 ≈ %.5f\n", result)
	}

	if _, err := c.Divide(1, 0); err != nil {
		fmt.Println("Помилка через інтерфейс:", err)
	}
}

func main() {
	fmt.Println("[main] Функція main стартувала")

	task1()
	task2()
}
