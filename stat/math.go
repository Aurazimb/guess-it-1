package stat

import (
	"math"
	"sort"
)

func Average(slice []float64) float64 {
	sum := 0.0
	for _, num := range slice {
		sum += num
	}
	x := sum / float64(len(slice))
	return x
}

func Median(slice []float64) float64 {
	sort.Float64s(slice)
	median := float64(len(slice) / 2.0)
	if len(slice)%2 == 0 {
		median = (slice[len(slice)/2-1] + slice[len(slice)/2]) / 2.0
	}
	return float64(median)
}

func Variance(slice []float64) float64 {
	av := Average(slice)
	var variance float64
	for _, el := range slice {
		dif := el - av
		variance += dif * dif
	}
	return variance / float64(len(slice))
}

func StandDev(slice []float64) float64 {
	variance := Variance(slice)
	Standdev := math.Sqrt(variance)
	return Standdev
}

func Skew(data []float64) float64 {
	n := float64(len(data))
	mean := 0.0
	for _, x := range data {
		mean += x
	}
	mean /= n

	var sumDiff3, sumDiff2 float64
	for _, x := range data {
		diff := x - mean
		sumDiff2 += diff * diff
		sumDiff3 += diff * diff * diff
	}

	stdDev := math.Sqrt(sumDiff2 / n)
	skew := sumDiff3 / (n * stdDev * stdDev * stdDev)

	return skew
}

func FindMode(data []float64) float64 {
	frequency := make(map[float64]float64)
	maxFrequency := 0.0
	var mode float64

	for _, value := range data {
		frequency[value]++
		if frequency[value] > maxFrequency {
			maxFrequency = frequency[value]
			mode = value
		}
	}
	return mode
}
