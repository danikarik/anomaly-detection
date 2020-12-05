package anomaly

import "math"

func calcMean(in []float64) float64 {
	var sum float64
	for _, v := range in {
		sum += v
	}

	return sum / float64(len(in))
}

func calcSD(in []float64) float64 {
	mean := calcMean(in)

	var sum float64
	for _, v := range in {
		sum += math.Pow(v-mean, 2)
	}

	return math.Sqrt(sum / float64(len(in)))
}
