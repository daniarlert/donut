package main

import (
	"fmt"
	"math"
	"time"
)

const (
	tau          = 2 * math.Pi
	thetaSpacing = 0.07
	phiSpacing   = 0.02
)

func reset[E float64 | byte](s []E, v E) {
	for i := range s {
		s[i] = v
	}
}

func donut() {
	A := 0.0
	B := 0.0

	z := make([]float64, 1760)
	b := make([]byte, 1760)

	fmt.Print("\033[H\033[2J")

	for {
		sA, cA := math.Sin(A), math.Cos(A)
		sB, cB := math.Sin(B), math.Cos(B)

		reset(b, ' ')
		reset(z, 0)

		for theta := 0.0; theta < tau; theta += thetaSpacing {
			sT, cT := math.Sin(theta), math.Cos(theta)

			for phi := 0.0; phi < tau; phi += phiSpacing {
				sP, cP := math.Sin(phi), math.Cos(phi)

				h := cT + 2
				D := 1 / (sP*h*sA + sT*cA + 5)
				t := sP*h*cA - sT*sA
				x := int(40 + 30*D*(cP*h*cB-t*sB))
				y := int(12 + 15*D*(cP*h*sB+t*cB))
				o := int(x + 80*y)

				N := int(8 * ((sT*sA-sP*cT*cA)*cB - sP*cT*sA - sT*cA - cP*cT*sB))
				if y < 22 && y > 0 && x > 0 && x < 80 && D > z[o] {
					z[o] = D

					point := 0
					if N > 0 {
						point = N
					}

					b[o] = ".,-~:;=!*#$@"[point]
				}
			}
		}

		print("\x1b[H")

		for k := 0; k < len(b); k++ {
			v := "\n"

			if k%80 > 0 {
				v = string(b[k])
			}

			fmt.Printf(v)

			A += 0.00004
			B += 0.00002
		}

		time.Sleep(16 * time.Millisecond)
	}
}

func main() {
	donut()
}
