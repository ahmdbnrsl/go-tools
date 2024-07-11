package maths

import "math"

func Pythagoras(alas, sisi int) any {
    return math.Sqrt(math.Pow(float64(alas), 2) + math.Pow(float64(sisi), 2))
}