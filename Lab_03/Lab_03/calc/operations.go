package calc

import (
	"errors"
	"fmt"
	"math"
)

func init() {
	fmt.Println("[calc] Пакет calc ініціалізовано (init викликано автоматично)")
}

func Sum(nums ...float64) float64 {
	total := 0.0
	for _, n := range nums {
		total += n
	}
	return total
}

func Max(nums ...float64) float64 {
	if len(nums) == 0 {
		return math.NaN()
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func Min(nums ...float64) float64 {
	if len(nums) == 0 {
		return math.NaN()
	}
	min := nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("ділення на нуль заборонено")
	}
	return a / b, nil
}
