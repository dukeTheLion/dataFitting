package main

import "fmt"

func sum(v []float64) float64 {
	s := 0.0

	for _, i := range v {
		s += i
	}

	return s
}

func quadMin(x []float64, y []float64) {
	upr := 0.0
	lwr := 0.0

	xSum := sum(x)
	ySum := sum(y)

	var b1 float64
	var b0 float64
	var xMean float64
	var yMean float64
	var r float64

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

	for i := 0; i < int(n); i++ {
		fmt.Printf("%10.3f %10.3f %10.3f %10.3f %10.3f %10.3f %10.3f %10.3f\n", y[i], x[i], xy[i], xx[i], yy[i], yc[i], e[i], ee[i])
	}

	fmt.Printf("%10.3f %10.3f %10.3f %10.3f %10.3f %10.3f %10.3f %10.3f\n", sum(y), sum(x), sum(xy), sum(xx), sum(yy), sum(yc), sum(e), sum(ee))
	fmt.Printf("b1=%7.4f b0=%7.4f n=%7.4f x|=%7.4f y|=%7.4f up=%7.4f lr=%7.4f R2=%7.4f", b1, b0, n, xMean, yMean, upr, lwr, r)

}

func main() {
	y := []float64{82, 91, 100, 68, 87, 73, 78, 80, 65, 84, 116, 76, 97, 100, 105, 77, 73, 78}
	x := []float64{71, 64, 43, 67, 56, 73, 68, 56, 76, 65, 45, 58, 45, 53, 49, 78, 73, 68}

	quadMin(x, y)
}
