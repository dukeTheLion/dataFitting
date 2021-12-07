package main

import (
	"dataFitting/Matrix"
	"fmt"
	"os"
)

func sum(v []float64) float64 {
	s := 0.0

	for _, i := range v {
		s += i
	}

	return s
}

func quadMin(x []float64, y []float64, sOutput bool) {
	upr := 0.0
	lwr := 0.0

	xSum := sum(x)
	ySum := sum(y)

	var r float64
	var b1 float64
	var b0 float64
	var xMean float64
	var yMean float64
	var out string

	n := float64(len(x))

	xy := make([]float64, 0, int(n))
	xx := make([]float64, 0, int(n))
	yy := make([]float64, 0, int(n))
	yc := make([]float64, 0, int(n))
	e := make([]float64, 0, int(n))
	ee := make([]float64, 0, int(n))

	for i := 0; i < int(n); i++ {
		xy = append(xy, x[i]*y[i])
		xx = append(xx, x[i]*x[i])
		yy = append(yy, y[i]*y[i])
	}

	xMean = xSum / n
	yMean = ySum / n
	b1 = (n*sum(xy) - xSum*ySum) / (n*sum(xx) - xSum*xSum)
	b0 = yMean - b1*xMean

	for i := 0; i < int(n); i++ {
		yc = append(yc, b0+b1*x[i])
		e = append(e, y[i]-yc[i])
		ee = append(ee, e[i]*e[i])
		upr += (yc[i] - yMean) * (yc[i] - yMean)
		lwr += (y[i] - yMean) * (y[i] - yMean)
	}

	r = upr / lwr

	out = "     ┌────────────┬────────────┬────────────┬────────────┬────────────┬────────────┬────────────┬────────────┐\n" +
		"     |     y      |      x     |     xy     |      x²    |      y²    |     y^     |      e     |      e²    |\n" +
		"     ├────────────┼────────────┼────────────┼────────────┼────────────┼────────────┼────────────┼────────────┤\n"

	for i := 0; i < int(n); i++ {
		out += fmt.Sprintf(" %03d | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f |\n",
			i, y[i], x[i], xy[i], xx[i], yy[i], yc[i], e[i], ee[i])
	}

	out += "     ├────────────┼────────────┼────────────┼────────────┼────────────┼────────────┼────────────┼────────────┤\n"

	out += fmt.Sprintf(" SUM | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f | %10.3f |\n",
		sum(y), sum(x), sum(xy), sum(xx), sum(yy), sum(yc), sum(e), sum(ee))

	out += "     └────────────┴────────────┴────────────┴────────────┴────────────┴────────────┴────────────┴────────────┘\n\n"

	out += fmt.Sprintf(" Number of lines ─► %.0f\n", n)
	out += fmt.Sprintf(" Mean of x ───────► %7.4f\n", xMean)
	out += fmt.Sprintf(" Mean of y ───────► %7.4f\n", yMean)
	out += fmt.Sprintf(" R ───────────────► %.4f ÷ %.4f = %7.4f\n", upr, lwr, r)
	out += fmt.Sprintf(" Func ────────────► f(x) = %.3fx%+.3f \n", b1, b0)

	outputFile, _ := os.Create("Output.txt")
	outputFile.WriteString(out)

	if sOutput {
		fmt.Print(out)
	}
}

func main() {
	//x := []float64{4.8, 7.7, 10.6, 12.5, 15.4, 18.2, 20.2, 25}
	//y := []float64{26.5, 26.8, 27.1, 27.3, 27.6, 27.9, 28.1, 28.6}

	inverse, err := Matrix.Inverse([][]float64{{2.0, 1, 1, 0},
		{4, 3, 3, 1},
		{8, 7, 9, 5},
		{6, 7, 9, 8}})

	if err != nil {
		return
	}

	fmt.Printf("%v", inverse)

	//quadMin(x, y, false)
}
