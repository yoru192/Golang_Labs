package main

import (
	"fmt"
	"math"
)

// ==================== ІНТЕРФЕЙС ====================

type Shape interface {
	Area() float64
	Perimeter() float64
}

// ==================== СТРУКТУРИ ====================

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// ==================== ДОПОМІЖНА ФУНКЦІЯ ====================

func printShapeInfo(name string, s Shape) {
	fmt.Printf("%s:\n", name)
	fmt.Printf("  Площа:    %.4f\n", s.Area())
	fmt.Printf("  Периметр: %.4f\n", s.Perimeter())
}

// ==================== MAIN ====================

func main() {
	// ---------- Частина 1: Масиви та слайси ----------

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := []int{5, 3, 8, 1, 9, 2, 7, 4, 6, 10}
	result := make([]int, 10)

	for i := 0; i < len(a); i++ {
		result[i] = a[i] + b[i]
	}

	fmt.Println("=== Масиви та слайси ===")
	fmt.Println("Масив a:  ", a)
	fmt.Println("Слайс b:  ", b)
	fmt.Println("Слайс result:", result)

	// ---------- Частина 2: Структури та інтерфейси ----------

	fmt.Println("\n=== Структури та інтерфейси ===")

	shapes := []struct {
		name  string
		shape Shape
	}{
		{"Circle (r=5)", Circle{Radius: 5}},
		{"Rectangle (4x6)", Rectangle{Width: 4, Height: 6}},
		{"Triangle (3, 4, 5)", Triangle{A: 3, B: 4, C: 5}},
	}

	for _, entry := range shapes {
		printShapeInfo(entry.name, entry.shape)
	}
}
