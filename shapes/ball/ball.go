// ball/ball.go
package ball

import (
    "math"
    "github.com/baduthutan/Popular-Programming-Language/shapes"
)

type Ball struct {
    Radius float64
}

func (b Ball) Volume() float64 {
    return (4.0 / 3.0) * math.Pi * math.Pow(b.Radius, 3)
}

func (b Ball) Area() float64 {
    return 4 * math.Pi * math.Pow(b.Radius, 2)
}

func NewBall(radius float64) shapes.Shape {
    return Ball{Radius: radius}
}
