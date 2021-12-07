package Matrix

import (
	"errors"
	"fmt"
)

func Inverse(k [][]float64) ([][]float64, error) {
	if len(k) != len(k[0]) {
		return [][]float64{}, errors.New(fmt.Sprintf("Matrix %dx%d is not valid", len(k), len(k[0])))
	}

	n := len(k)
	invrs := make([][]float64, 0, n)

	for i := 0; i < n; i++ {
		val := make([]float64, 0, n)
		for j := 0; j < n; j++ {
			if j == i {
				val = append(val, 1)
			} else {
				val = append(val, 0)
			}
		}
		invrs = append(invrs, val)
	}

	for i := 0; i < n-1; i++ {
		for j := i; j < n-1; j++ {
			op := k[j+1][i] / k[i][i]

			if k[j+1][i] != 0 {
				for w := 0; w < n; w++ {
					k[i][w] = k[i][w] * op
					invrs[i][w] = invrs[i][w] * op
					k[j+1][w] = k[i][w] - k[j+1][w]
					invrs[j+1][w] = invrs[i][w] - invrs[j+1][w]
				}
			}
		}
	}

	for i := n - 1; i > 0; i-- {
		for j := i; j > 0; j-- {
			op := k[j-1][i] / k[i][i]

			if k[j-1][i] != 0 {
				for w := 0; w < n; w++ {
					k[i][w] = k[i][w] * op
					invrs[i][w] = invrs[i][w] * op
					k[j-1][w] = k[i][w] - k[j-1][w]
					invrs[j-1][w] = invrs[i][w] - invrs[j-1][w]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		op := k[i][i]

		for j := 0; j < n; j++ {
			invrs[i][j] = invrs[i][j] / op
			k[i][j] = k[i][j] / op
		}
	}

	return invrs, nil
}
