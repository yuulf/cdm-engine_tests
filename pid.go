package main

import (
	"bytes"
	"github.com/wcharczuk/go-chart"
	"io/ioutil"
	"math"
	"os"
	"path"
)

// NewPIDController returns a new PIDController using the given gain values.
func NewPIDController(p, i, d float64) *PIDController {
	return &PIDController{p: p, i: i, d: d, outMin: math.Inf(-1), outMax: math.Inf(0)}
}

// PIDController implements a PID controller.
type PIDController struct {
	p         float64 // proportional gain
	i         float64 // integral gain
	d         float64 // derrivate gain
	setpoint  float64 // current setpoint
	prevValue float64 // last process value
	integral  float64 // integral sum
	outMin    float64 // Output Min
	outMax    float64 // Output Max
}

func (data *PIDController) pidCalculate(setpoint float64, pv float64) float64 {
	// Error
	err := setpoint - pv

	// Proportional term
	pout := data.p * err

	// Integral term
	data.integral += err
	iout := data.i * data.integral

	// Derivative term
	derivative := err - data.prevValue
	dout := data.d * derivative

	// Total output
	output := pout + iout + dout

	// Restrict
	if output > data.outMax {
		output = data.outMax
	}
	if output < data.outMin {
		output = data.outMin
	}
	data.prevValue = err

	return output
}

func main() {
	p := NewPIDController(1, 0, 0.4)
	p.setpoint = 0
	p.outMax = 3000
	p.outMin = -3000
	val := 180.0
	xValues := []float64{}
	yValues := []float64{}
	forceValues := []float64{}

	currentSpeed := 0.0

	for i := 0.0; i < 100; i += 0.1 {
		yValues = append(yValues, val)
		inc := p.pidCalculate(0, val)
		xValues = append(xValues, float64(i))
		if currentSpeed > 100.0 {
			currentSpeed -= 100
		}
		if currentSpeed < (-100.0) {
			currentSpeed += 100
		}
		currentSpeed += inc
		forceValues = append(forceValues, inc)

		val += currentSpeed * 0.1

	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
			chart.ContinuousSeries{
				Name:    "Скорость",
				YAxis:   chart.YAxisSecondary,
				XValues: xValues,
				YValues: forceValues,
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	saveTo := path.Join(os.Getenv("HOME"), "1.png")
	ioutil.WriteFile(saveTo, buffer.Bytes(), 0644)
}
