package mymath

import (
	"math"
)

// Abs возвращает абсолютное значение числа.
func Abs(x float64) float64 {
	return math.Abs(x)
}

// Sqrt возвращает квадратный корень числа.
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Pow возвращает x в степени y.
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Max возвращает максимальное значение из двух чисел.
func Max(x, y float64) float64 { return math.Max(x, y) }

//{
//	if x > y {
//		return x
//	}
//	return y
//}

// Min возвращает минимальное значение из двух чисел.
func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

// Clamp возвращает значение, ограниченное минимальным и максимальным значениями.
func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

// IsNaN проверяет, является ли число "не числом" (Not-a-Number).
func IsNaN(x float64) bool {
	return math.IsNaN(x)
}

// IsInf проверяет, является ли число бесконечным.
func IsInf(x float64) bool {
	return math.IsInf(x, 0)
}

// Round возвращает число, округленное до ближайшего целого.
func Round(x float64) float64 {
	return math.Round(x)
}

// Trunc возвращает целую часть числа.
func Trunc(x float64) float64 {
	return math.Trunc(x)
}

// Floor возвращает наибольшее целое число, не превосходящее x.
func Floor(x float64) float64 {
	return math.Floor(x)
}

// Ceil возвращает наименьшее целое число, не меньшее x.
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Nearest возвращает ближайшее целое число к x.
func Nearest(x float64) float64 {
	return Round(x)
}

// Float64sEqual проверяет, равны ли два числа с плавающей точкой.
func Float64sEqual(x, y float64) bool {
	return math.Abs(x-y) < 1e-1
}
func Yn(n int, x float64) float64 { return math.Yn(n, x) }

func Acos(x float64) float64 {
	return math.Acos(x)
}
func Acosh(x float64) float64 {
	return math.Acosh(x)
}
func Asin(x float64) float64 {
	return math.Asin(x)
}
