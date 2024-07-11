package maths

import "math"

type VolumeType struct{}

var Volume = VolumeType{}

func (_ VolumeType) Kubus(sisi int) any {
    return math.Pow(float64(sisi), 3)
}

func (_ VolumeType) Balok(panjang, lebar, tinggi int) any {
    return panjang * lebar * tinggi
}

func (_ VolumeType) Bola(radius int) any {
    if radius % 7 != 0 {
        return 3.14 * math.Pow(float64(radius), 3)
    } else {
        return float64(22) / 7 * math.Pow(float64(radius), 3)
    }
}