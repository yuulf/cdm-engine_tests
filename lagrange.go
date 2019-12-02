package main

func InterpolateLagrangePolynomial(x float64, xValues []float64, yValues []float64, size int) float64 {
	lagrangePol := 0.0
	for index := 0; index < size; index++ {
		basicsPol := 1.0
		for index2 := 0; index2 < size; index2++ {
			if index != index2 {
				basicsPol *= (x - xValues[index2]) / (xValues[index] - xValues[index2])
			}
		}
		lagrangePol += basicsPol * yValues[index]
	}
	return lagrangePol
}

func TestF(x float64) float64 {
	return x*x*x + 3*x*x + 3*x + 1

}

func main() {
	const size = 10
	xValues := []float64{}
	yValues := []float64{}
	for index := 0; index < size; index++ {
		xValues = append(xValues, float64(index))
		yValues = append(yValues, TestF(float64(index)))
	}
	println(InterpolateLagrangePolynomial(15, xValues, yValues, size))
}
